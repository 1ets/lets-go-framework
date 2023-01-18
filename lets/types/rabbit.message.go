package types

type RabbitBody struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

type ReplyTo struct {
	Exchange      string `json:"exchange"`
	RoutingKey    string `json:"routing_key"`
	CorrelationId string `json:"correlation_id"`
}
