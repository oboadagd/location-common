package dto

import "time"

type LocationHistory struct {
	tableName struct{}  `pg:"location_history,alias:locationHistory"`
	Id        int64     `json:"id" pg:",pk"`
	UserName  string    `json:"userName"  pg:"username, unique, notnull, alias:userName"`
	Latitude  float64   `json:"latitude"  pg:", notnull"`
	Longitude float64   `json:"longitude" pg:", notnull"`
	UpdatedAt time.Time `json:"updatedAt" pg:"updated_at, notnull, alias:updatedAt"`
	Distance  float64   `json:"distance"  pg:", notnull"`
}
