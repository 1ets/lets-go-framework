package data

type RequestExample struct {
	Name string `json:"name"`
}

type ResponseExample struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
