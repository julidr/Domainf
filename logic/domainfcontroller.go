package logic

import (
	"Domainf/models"
	"Domainf/repository"
	"encoding/json"
	"fmt"
	"github.com/likexian/whois-go"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var connection = repository.GetConnection("root", "Juli", "26257", "domainf", "disable")

// Handle all the logic to retrieve the information of the specified domain.
// Create or update the domain info in the database, build the response information
func GetServerInformation(host string) {
	var servers []string
	result := &models.Information{}
	analyze := callAnalyzeEndpoint(host)
	endpointsCount := len(analyze.Endpoints)
	result.Servers = make([]models.Server, 0)
	for i := 0; i < endpointsCount; i++ {
		server := models.Server{}
		server.SetAddress(analyze.Endpoints[i].IpAddress)
		server.SetSslGrade(analyze.Endpoints[i].Grade)
		extraInfo := getOwnerAndCountry(analyze.Endpoints[i].IpAddress)
		server.SetOwner(extraInfo[0])
		server.SetCountry(extraInfo[1])
		servers = append(servers, analyze.Endpoints[i].IpAddress)
		result.Servers = append(result.Servers, server)
	}
	lowerGrade := calculateLowerSslGrade(analyze.Endpoints)
	result.SetSslGrade(lowerGrade)
	domain := getDomain(host)
	if domain.Id() == "" {
		createDomain(host, lowerGrade, servers)
	} else {
		serversChanged := compareServers(servers, domain.Servers())
		result.SetServersChanged(serversChanged)
		if serversChanged == true {
			domain.SetServers(servers)
		}
		result.SetPreviousSslGrade(domain.SslGrade())
	}
	domain.SetSslGrade(lowerGrade)
	domain.SetUpdatedAt(time.Now())
	updateDomain(domain)
	response, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Something failed with the response: ", err)
	}
	fmt.Println(string(response))
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
		log.Fatal("Something failed while getting whois information: ", err)
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
	return sslGrades[lowerGrade]
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