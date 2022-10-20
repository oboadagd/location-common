// Package dto implements definition of structs to
// store objects that represent database entities
// and http requests and responses
package dto

import "time"

// Location describes a database Location entity. Defines a geographic
// point of a username at a given date through latitude and longitude coordinates.
type Location struct {
	tableName struct{}  `pg:"location,alias:location"`                                   // name of the table. Control field not visible
	Id        int64     `json:"id" pg:",pk"`                                             // record identifier
	UserName  string    `json:"userName" pg:"username, notnull, unique, alias:userName"` // username
	Latitude  float64   `json:"latitude" pg:",use_zero, notnull"`                        // latitude coordinate of a geographic point
	Longitude float64   `json:"longitude" pg:",use_zero, notnull"`                       // longitude coordinate of a geographic point
	UpdatedAt time.Time `json:"updatedAt" pg:"updated_at, notnull, alias:updatedAt"`     // date of later update
}
