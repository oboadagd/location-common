package dto

import "time"

// LocationHistory describes a database LocationHistory entity. Defines a historic
// register of geographic points of usernames by dates through latitude and longitude
// coordinates.
type LocationHistory struct {
	tableName struct{}  `pg:"location_history,alias:locationHistory"`                // name of the table. Control field not visible
	Id        int64     `json:"id" pg:",pk"`                                         // record identifier
	UserName  string    `json:"userName"  pg:"username, notnull, alias:userName"`    // username
	Latitude  float64   `json:"latitude"  pg:",use_zero, notnull"`                   // latitude coordinate of a geographic point
	Longitude float64   `json:"longitude" pg:",use_zero, notnull"`                   // longitude coordinate of a geographic point
	UpdatedAt time.Time `json:"updatedAt" pg:"updated_at, notnull, alias:updatedAt"` // date of update
	Distance  float64   `json:"distance" pg:",use_zero"`                             // traveled distance by username from last to current location
}
