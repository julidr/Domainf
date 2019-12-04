package logic

import (
	"Domainf/models"
	"Domainf/repository"
	"encoding/json"
	"fmt"
	"github.com/likexian/whois-go"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// TODO handle this with secrets
var connection = repository.GetConnection("root", "Juli", "26257", "domainf", "disable")

// Handle all the logic to retrieve the information of the specified domain.
// Create or update the domain info in the database, build the response information
func GetServerInformation(host string) []byte {
	var servers []string
	result := &models.Information{}
	// Call the SSL Labs endpoint to get the information of the domain
	analyze := callAnalyzeEndpoint(host)
	endpointsCount := len(analyze.Endpoints)
	result.Servers = make([]models.Server, 0)
	// Iterate over the returned endpoints to get their country and owner
	for i := 0; i < endpointsCount; i++ {
		server := models.Server{}
		server.SetAddress(analyze.Endpoints[i].IpAddress)
		server.SetSslGrade(analyze.Endpoints[i].Grade)
		// Call Whois library to get Owner and Country
		extraInfo := getOwnerAndCountry(analyze.Endpoints[i].IpAddress)
		server.SetOwner(extraInfo[0])
		server.SetCountry(extraInfo[1])
		servers = append(servers, analyze.Endpoints[i].IpAddress)
		result.Servers = append(result.Servers, server)
	}
	// If the SSL Labs didn't return an array of servers
	// Then this can be considered like the is down (not really)
	if len(result.Servers) == 0 {
		result.SetIsDown(true)
	}
	// Calculate which one is the lowest grade in the array of servers
	lowerGrade := calculateLowerSslGrade(analyze.Endpoints)
	result.SetSslGrade(lowerGrade)
	// Get the website of the domain and try to get its logo and title
	headInformation := getLogoAndTitle(analyze.Host, analyze.Protocol)
	result.SetLogo(headInformation[0])
	result.SetTitle(headInformation[1])
	// Search in the database if a domain related to that host already exist
	domain := getDomain(host)
	// Create a new domain in the database in case that didn't exist
	if domain.Id() == "" {
		createDomain(host, lowerGrade, servers)
	} else {
		// Compare if the returned servers are different from the previous stored servers
		serversChanged := compareServers(servers, domain.Servers())
		result.SetServersChanged(serversChanged)
		if serversChanged == true {
			domain.SetServers(servers)
		}
		result.SetPreviousSslGrade(domain.SslGrade())
	}
	// Update the domain with its new grade and time
	domain.SetSslGrade(lowerGrade)
	domain.SetUpdatedAt(time.Now())
	updateDomain(domain)
	// Transform the Information Struct to a JSON
	response, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Something failed with the response: ", err)
	}
	return response
}

// Call the database to retrieve the history of domains
func GetServersHistory() []byte {
	domainsHistory := repository.GetDomainsHistory(connection)
	history := models.History{domainsHistory}
	response, err := json.Marshal(history)
	if err != nil {
		log.Fatal("Something failed with the response: ", err)
	}
	return response
}

// Make a GET request to the SSL Labs API and retrieve the analyzed information for a host.
func callAnalyzeEndpoint(host string) models.Analyze {
	var analyze models.Analyze
	url := fmt.Sprintf("https://api.ssllabs.com/api/v3/analyze?host=%v", host)
	response, requestError := http.Get(url)
	if requestError != nil {
		log.Fatal("Something failed with the request to the servers info: ", requestError)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &analyze)
	return analyze
}

// Make a whois call to the specified IP to retrieve information like the owner and country of the server
func getOwnerAndCountry(ip string) [2]string {
	var owner string
	var country string
	var information [2]string
	whoIsResult, err := whois.Whois(ip, "whois.arin.net")
	if err != nil {
		log.Println("Something failed while getting whois information: ", err)
	}
	lines := strings.Split(whoIsResult, "\n")
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "Organization:") {
			owner = strings.TrimSpace(strings.Split(lines[i], ":")[1])
		}
		if strings.Contains(lines[i], "Country:") {
			country = strings.TrimSpace(strings.Split(lines[i], ":")[1])
		}
	}
	information[0] = owner
	information[1] = country
	return information
}

// Calculate the lowest ssl grade for a given list of possible grades
func calculateLowerSslGrade(endpoints []models.Endpoint) string {
	var grade string
	sslGrades := [7]string{"F", "E", "D", "C", "B", "A", "A+"}
	lowerGrade := 10
	gradesCount := len(sslGrades)
	endpointsCount := len(endpoints)
	for i := 0; i < endpointsCount; i++ {
		grade := endpoints[i].Grade
		for j := 0; j < gradesCount; j++ {
			if grade == sslGrades[j] {
				if j < lowerGrade {
					lowerGrade = j
				}
			}
		}
	}
	if lowerGrade >= 0 && lowerGrade <= 6 {
		grade = sslGrades[lowerGrade]
	}
	return grade
}

// Return the domain for given host in the database
func getDomain(host string) models.Domain {
	domain := repository.GetDomainByHost(connection, host)
	return domain
}

// Create a domain in the database
func createDomain(host string, sslGrade string, endpoints []string) {
	domain := models.Domain{}
	domain.SetHost(host)
	domain.SetSslGrade(sslGrade)
	domain.SetServers(endpoints)
	domain.SetCreatedAt(time.Now())
	domain.SetUpdatedAt(time.Now())
	repository.CreateDomain(connection, domain)
}

// Update the domain in the database with the given information
func updateDomain(domain models.Domain) {
	repository.UpdateDomain(connection, domain)
}

// Compare if two servers and verify if they changed or not
func compareServers(requestServers []string, databaseServers[]string) bool {
	serverChanged := false
	if len(requestServers) != len(databaseServers) {
		serverChanged = true
	} else {
		for i := 0; i < len(requestServers); i++ {
			if requestServers[i] != databaseServers[i] {
				serverChanged = true
			}
		}
	}
	return serverChanged
}

// Call the possible web page of a domain and extract its logo and title from the HTML content
func getLogoAndTitle(host string, protocol string) [2]string {
	var logo string
	var title string
	var headInformation [2]string
	url := fmt.Sprintf("%v://%v", protocol, host)
	response, requestError := http.Get(url)
	if requestError != nil {
		log.Println("Something failed with the request to the servers info: ", requestError)
	}
	defer response.Body.Close()
	htmlResponse, _ := ioutil.ReadAll(response.Body)
	pageContent := string(htmlResponse)
	// Get only the content of the HEAD tag
	headStartIndex := strings.Index(pageContent, "<head>")
	headEndIndex := strings.Index(pageContent, "</head>")
	headContent := pageContent[headStartIndex:headEndIndex]
	headDoc, err := html.Parse(strings.NewReader(headContent))
	if err != nil {
		log.Println("something failed while parsing the HTML", err)
	}
	logo = searchIcon(headDoc)
	title = searchTitle(headDoc)
	headInformation[0] = logo
	headInformation[1] = title
	return headInformation
}

// Given an HTML node search for the shortcut icon from its head
func searchIcon(n *html.Node) string {
	var logo string
	if n.Type == html.ElementNode && n.Data == "link" {
		for i:= 0; i < len(n.Attr); i++ {
			if strings.Contains(n.Attr[i].Val, "shortcut icon") {
				for j := 0; j < len(n.Attr); j++ {
					if n.Attr[j].Key == "href" {
						return n.Attr[j].Val
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		logo = searchIcon(c)
		if logo != "" {
			return logo
		}
	}
	return logo
}

// Given an HTML node search for the title of the Head
func searchTitle (n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = searchTitle(c)
		if title != "" {
			return title
		}
	}
	return title
}