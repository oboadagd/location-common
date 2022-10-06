package dto

import "time"

type LocationHistory struct {
	tableName struct{}  `pg:"location_history,alias:locationHistory"`
	Id        int64     `json:"id" pg:",pk"`
	UserName  string    `json:"userName" pg:"username, alias:userName"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	UpdatedAt time.Time `json:"updatedAt" pg:"updated_at, alias:updatedAt"`
	Distance  float64   `json:"distance"`
}
