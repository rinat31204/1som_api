package service

type EventBus interface {
	Publish(eventName string, event any)
	Subscribe(eventName string, handler func(any))
}
