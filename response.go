package goshopee

type BaseResponse struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Warning   string `json:"warning,omitempty"`
}
