package dto

// GetPagesUsersByLocationAndRadiusResponse is a http response of GetUsersByLocationAndRadius service
type GetUsersByLocationAndRadiusResponse struct {
	Users      []Location `json:"users"`      // username's location list located in a given radius
	TotalItems uint64     `json:"totalItems"` // total number of items
	TotalPages uint64     `json:"totalPages"` // total number of pages
}
