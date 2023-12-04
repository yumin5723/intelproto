package events

type OrderCreatedEvent struct {
	OrderId string
	UserId  string
	Amount  int
}
