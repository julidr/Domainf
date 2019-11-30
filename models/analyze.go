package models

type Analyze struct {
	Host string `json:"host"`
	Endpoints []struct {
		IpAddress string `json:"ipAddress"`
		Grade string `json:"grade"`
	} `json:"endpoints"`
}