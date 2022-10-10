package dto

import "time"

type Location struct {
	tableName struct{}  `pg:"location,alias:location"`
	Id        int64     `json:"id" pg:",pk"`
	UserName  string    `json:"userName" pg:"username, notnull, unique, alias:userName"`
	Latitude  float64   `json:"latitude" pg:", notnull"`
	Longitude float64   `json:"longitude" pg:", notnull"`
	UpdatedAt time.Time `json:"updatedAt" pg:"updated_at, notnull, alias:updatedAt"`
}
