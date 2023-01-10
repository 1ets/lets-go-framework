package structs

type HttpAccountRequestFind struct {
	Id int32 `uri:"id"`
}

type HttpAccountRequestRegister struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type HttpAccountRequestUpdate HttpAccountRequestRegister

type HttpAccountResponseDefault struct {
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type DefaultHttpResponse struct {
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
