package response

type ErrorResponse struct {
	FailedFields []ValidationError `json:"failed_fields"`
}

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
	Param string `json:"param,omitempty"`
}
