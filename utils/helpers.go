package utils

import "strings"

// Transform a slice to the specific ARRAY format of the DB
func FormatServersToDBArray(servers []string) string {
	formatArray := "ARRAY["
	serversCount := len(servers)
	for i := 0; i < serversCount; i++ {
		if i != len(servers)-1 {
			formatArray += "'" + servers[i] + "',"
		} else {
			formatArray += "'" + servers[i] + "'"
		}
	}
	formatArray += "]"
	return formatArray
}

// Transform a string ARRAY from the DB to a slice
func FormatDBArrayToServers(dbArray string) []string {
	dbArray = strings.Replace(dbArray, "{", "", -1)
	dbArray = strings.Replace(dbArray, "}", "", -1)
	return strings.Split(dbArray, ",")
}
