// Contains all necessary functions to connect with the CockroachDB
package logic

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// GetConnection returns the connection CockroachDB logic based on the given information
func GetConnection(user string, host string, port string, sslMode bool) *sql.DB {
	dataSourceName := fmt.Sprintf("postgresql://%v@%v:%v?sslmode=%v", user, host, port, sslMode)
	connection, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return connection
}