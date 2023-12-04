package events

const ORDER_CREATED_EVENT_TOPIC = "order.created"

type OrderCreatedEvent struct {
	OrderId string
	UserId  string
	Amount  int
}
