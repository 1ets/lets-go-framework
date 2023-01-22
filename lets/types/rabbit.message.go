package types

type RabbitBody struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
