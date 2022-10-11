package dto

type GetLastByUserNameResponse struct {
	Username  string  `json:"userName"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
