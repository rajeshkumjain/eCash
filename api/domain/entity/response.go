package entity

// ResponseResult : not clear where to use - will plan shortly
type ResponseResult struct {
	Error string `json:"error, omitempty"`
	Success string `json:"success, omitempty"`
	Code   string `json:"code, omitempty"`
	Message string `json:"message, omitempty"`
}
