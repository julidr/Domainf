package main

import (
	"Domainf/logic"
	"fmt"
)

func main() {
	// Connect to the "domainf" repository.
	fmt.Println("Connecting")
	//connection := repository.GetConnection("root", "Juli", "26257", "domainf", "disable")
	//myDomain := models.Domain{}
	//myDomain.SetHost("google.com")
	//myDomain.SetServers([]string{"172.217.5.110", "2607:f8b0:4005:808:0:0:0:200e"})
	//myDomain.SetSslGrade("A")
	//myDomain.SetCreatedAt(time.Now())
	//myDomain.SetUpdatedAt(time.Now())
	//repository.CreateDomain(connection, myDomain)
	//domain := repository.GetDomainByHost(connection, "google.com")
	//if domain.Id() != "" {
	//	fmt.Println(domain)
	//	fmt.Println(domain.CreatedAt().Format("2006-01-02 15:04:05"))
	//}
	//domain.SetServers([]string{"172.217.6.46", "2607:f8b0:4005:808:0:0:0:200e"})
	//domain.SetUpdatedAt(time.Now())
	//repository.UpdateDomain(connection, domain)
	//endpoints := []models.Endpoint{{"1", "A+"},{"1", "A"},{"1", "A"}}
	//lowest := logic.CalculateLowerSslGrade(endpoints)
	//fmt.Println(lowest)
	logic.GetServerInformation("google.com")
}
