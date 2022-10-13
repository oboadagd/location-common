package dto

// Response is a simple http response with only a string message.
type Response struct {
	Message string `json:"message"` // string message
}
