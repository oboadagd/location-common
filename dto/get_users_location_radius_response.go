package dto

// GetUsersByLocationAndRadiusResponse is a http response of GetUsersByLocationAndRadius service
type GetUsersByLocationAndRadiusResponse struct {
	Users []Location `json:"users"` // username's location list located in a given radius
}
