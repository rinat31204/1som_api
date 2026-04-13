package inmemory

import (
	"log/slog"
	"sync"
)

type EventBus struct {
	subscribers map[string][]func(any)
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(any)),
	}
}

func (b *EventBus) Subscribe(eventName string, handler func(any)) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscribers[eventName] = append(b.subscribers[eventName], handler)
}

func (b *EventBus) Publish(eventName string, event any) {
	b.mu.RLock()
	handlers := b.subscribers[eventName]
	b.mu.RUnlock()

	for _, h := range handlers {
		go func(handler func(any), ev any) {
			defer func() {
				if r := recover(); r != nil {
					slog.Error("handler error: ", r)
				}
			}()
			handler(ev)
		}(h, event)
	}
}
