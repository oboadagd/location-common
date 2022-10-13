package dto

// GetDistanceTraveledResponse is http response of GetDistanceTraveled service
type GetDistanceTraveledResponse struct {
	Username      string  `json:"userName"`      // username
	TotalDistance float64 `json:"totalDistance"` // accumulated total traveled distance
}
