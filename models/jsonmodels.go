package models

// Represent the JSON response of the Analyze endpoint into a struct to easy manipulation
type Analyze struct {
	Host string `json:"host"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	IpAddress string `json:"ipAddress"`
	Grade string `json:"grade"`
}

// Represent the JSON response of the information requested for a certain domain
type Information struct {
	Servers [] Server `json:"servers"`
	ServersChanged bool `json:"servers_changed"`
	SslGrade string `json:"ssl_grade"`
	PreviousSslGrade string `json:"previous_ssl_grade"`
	Logo string `json:"logo"`
	Title string `json:"title"`
	IsDown bool `json:"is_down"`
}

// Set the value of the lower SslGrade of all the returned servers
func (information *Information) SetSslGrade(sslGrade string) {
	information.SslGrade = sslGrade
}

// Set the value of true if the servers changed or false if they didn't
func (information *Information) SetServersChanged(serversChanged bool) {
	information.ServersChanged = serversChanged
}

// Set the value of PreviousSslGrade if it exist
func (information *Information) SetPreviousSslGrade(previousSslGrade string) {
	information.PreviousSslGrade = previousSslGrade
}

// Represent the servers of the JSON response that have the requested information of a certain domain
type Server struct {
	Address string `json:"address"`
	SslGrade string `json:"ssl_grade"`
	Country string `json:"country"`
	Owner string `json:"owner"`
}
// Set the value of the Address to the respective server in the array of servers
func (server *Server) SetAddress(address string) {
	server.Address = address
}

// Set the value of the SslGrade to the respective server in the array of servers
func (server *Server) SetSslGrade(sslGrade string) {
	server.SslGrade = sslGrade
}

// Set the value of the Country to the respective server in the array of servers
func (server *Server) SetCountry(country string) {
	server.Country = country
}

// Set the value of the Owner to the respective server in the array of servers
func (server *Server) SetOwner(owner string) {
	server.Owner = owner
}