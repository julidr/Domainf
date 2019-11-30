package main

import (
	"Domainf/logic"
	"Domainf/models"
	"fmt"
	"time"
)

func main() {
	// Connect to the "domainf" logic.
	fmt.Println("Connecting")
	connection := logic.GetConnection("root", "Juli", "26257", "domainf", "disable")
	myDomain := models.Domain{}
	myDomain.SetHost("google.com")
	myDomain.SetServers([]string{"172.217.5.110", "2607:f8b0:4005:808:0:0:0:200e"})
	myDomain.SetSslGrade("A")
	myDomain.SetCreatedAt(time.Now())
	myDomain.SetUpdatedAt(time.Now())
	//logic.CreateDomain(connection, myDomain)
	domain := logic.GetDomainByHost(connection, "google.com")
	fmt.Println(domain)
	fmt.Println(domain.CreatedAt().Format("2006-01-02 15:04:05"))
}
