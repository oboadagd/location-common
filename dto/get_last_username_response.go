package dto

// GetLastByUserNameResponse is a response of GetLastByUserName method
type GetLastByUserNameResponse struct {
	Username  string  `json:"userName"`  // found username
	Latitude  float64 `json:"latitude"`  // later latitude coordinate of found username
	Longitude float64 `json:"longitude"` // later longitude coordinate of found username
}
