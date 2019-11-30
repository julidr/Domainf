package logic

import (
	"Domainf/models"
	"encoding/json"
	"fmt"
	"github.com/likexian/whois-go"
	"io/ioutil"
	"log"
	"net/http"
)

// Handle all the logic to retrieve the information of the specified domain.
// Create or update the domain info in the database, build the response information
func GetServerInformation(host string) {
	analyze := callAnalyze(host)
	fmt.Println(analyze.Host)
	fmt.Println(analyze.Endpoints)
	endpointsCount := len(analyze.Endpoints)
	for i := 0; i < endpointsCount; i++ {
		fmt.Println(analyze.Endpoints[i])
		makeWhoIs(analyze.Endpoints[i].IpAddress)
	}
}

// Make a GET request to the SSL Labs API and retrieve the analyzed information for a host.
func callAnalyze(host string) models.Analyze {
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

func makeWhoIs(ip string) {
	whoisResult, err := whois.Whois(ip, "whois.arin.net")
	if err != nil {
		log.Fatal("Something failed while getting whois information: ", err)
	}
	fmt.Println(whoisResult)
}
