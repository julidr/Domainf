// Contain struct representations of the different database tables
package models

import "time"

// Representation of the domain table
type Domain struct {
	id        string
	host      string
	servers   []string
	sslGrade  string
	createdAt time.Time
	updatedAt time.Time
}

// Return the id of a domain
func (domain Domain) Id() string {
	return domain.id
}

// Set the if of a Domain
func (domain *Domain) SetId(id string) {
	domain.id = id
}

// Return the host value of a domain
func (domain Domain) Host() string {
	return domain.host
}

// Set the host value of a domain
func (domain *Domain) SetHost(host string) {
	domain.host = host
}

// Return the servers of a domain
func (domain Domain) Servers() []string {
	return domain.servers
}

// Set the array of servers ips of a domain
func (domain *Domain) SetServers(servers []string) {
	domain.servers = servers
}

// Return the ssl grade value of a domain
func (domain Domain) SslGrade() string {
	return domain.sslGrade
}

// Set the ssl grade of a domain
func (domain *Domain) SetSslGrade(sslGrade string) {
	domain.sslGrade = sslGrade
}

// Return the creation date of a domain
func (domain Domain) CreatedAt() time.Time {
	return domain.createdAt
}

// Set the creation date of a domain
func (domain *Domain) SetCreatedAt(createdAt time.Time) {
	domain.createdAt = createdAt
}

// Return the update date of a domain
func (domain Domain) UpdatedAt() time.Time {
	return domain.updatedAt
}

// Set the update date of a domain
func (domain *Domain) SetUpdatedAt(updatedAt time.Time) {
	domain.updatedAt = updatedAt
}

