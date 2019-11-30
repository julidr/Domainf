package logic

import (
	"Domainf/models"
	"Domainf/utils"
	"database/sql"
	"fmt"
	"log"
	"time"
)

var datetimeLayout string = "2006-01-02 15:04:05"
var timezoneLayout string = "2006-01-02T15:04:05Z"

// CreateDomain save in the database a Domain with its respective information
func CreateDomain(db *sql.DB, domain models.Domain) {
	formatArray := utils.FormatServersToDBArray(domain.Servers())
	query := fmt.Sprintf("INSERT INTO domain (host, servers, ssl_grade, created_at, updated_at) "+
		"VALUES ('%v', %v, '%v', '%v', '%v');", domain.Host(), formatArray, domain.SslGrade(),
		domain.CreatedAt().Format(datetimeLayout), domain.UpdatedAt().Format(datetimeLayout))
	if _, err := db.Exec(query); err != nil {
		log.Fatal("Something failed with the creation of the Domain: ", err)
	}
}

// Search and return a Domain based on the given host
func GetDomainByHost(db *sql.DB, host string) models.Domain {
	domain := models.Domain{}
	query := fmt.Sprintf("SELECT * FROM domain WHERE host = '%v';", host)
	rows, err := db.Query(query);
	if err != nil {
		log.Fatal("Something failed while searching for " + host + ": ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var host string
		var servers string
		var sslGrade string
		var createdAt string
		var updatedAt string
		if err := rows.Scan(&id, &host, &servers, &sslGrade, &createdAt, &updatedAt); err != nil {
			log.Fatal("Something failed with the Scan: ", err)
		}
		domain.SetId(id)
		domain.SetHost(host)
		formatServers := utils.FormatDBArrayToServers(servers)
		domain.SetServers(formatServers)
		domain.SetSslGrade(sslGrade)
		parseCreatedAt, _ := time.Parse(timezoneLayout, createdAt)
		domain.SetCreatedAt(parseCreatedAt)
		parseUpdatedAt, _ := time.Parse(timezoneLayout, updatedAt)
		domain.SetUpdatedAt(parseUpdatedAt)
	}
	return domain
}
