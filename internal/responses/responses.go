package responses

// ErrorResponse standard API error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse standard API success response
type SuccessResponse struct {
	Message string `json:"message"`
}
