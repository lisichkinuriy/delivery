package ports

type IBasketConfirmedConsumer interface {
	Consume() error
	Close() error
}
