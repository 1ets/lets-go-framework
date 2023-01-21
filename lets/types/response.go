package types

// Extender internal response
type Response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
