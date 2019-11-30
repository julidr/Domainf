// Contains all necessary functions to connect with the CockroachDB
package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// GetConnection returns the connection CockroachDB repository based on the given information
func GetConnection(user string, host string, port string, database string, sslMode string) *sql.DB {
	dataSourceName := fmt.Sprintf("postgresql://%v@%v:%v/%v?sslmode=%v", user, host, port, database, sslMode)
	connection, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("Error connection to the database: ", err)
	}
	return connection
}

// CloseConnection handle the closure of any given connection
func CloseConnection(db *sql.DB) {
	db.Close()
}
