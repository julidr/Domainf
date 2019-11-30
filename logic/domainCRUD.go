package logic

import (
	"Domainf/models"
	"database/sql"
	"fmt"
	"log"
)

var timeLayout string = "2006-01-02 15:04:05"

func CreateDomain(db *sql.DB, domain models.Domain) {
	formatArray := domain.FormatServersToDBArray()
	query := fmt.Sprintf("INSERT INTO domain (host, servers, ssl_grade, created_at, updated_at) "+
		"VALUES ('%v', %v, '%v', '%v', '%v');", domain.Host(), formatArray, domain.SslGrade(),
		domain.CreatedAt().Format(timeLayout), domain.UpdatedAt().Format(timeLayout))
	fmt.Println(query)
	if _, err := db.Exec(query); err != nil {
		log.Fatal("Something failed with the creation of the Domain: ", err)
	}
}
