package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockEventBus мок для EventBus (если нужно проверять вызовы)
type MockEventBus struct {
	mock.Mock
}

func (m *MockEventBus) Subscribe(event string, handler func(any)) {
	m.Called(event, handler)
}

func (m *MockEventBus) Publish(event string, data any) {
	m.Called(event, data)
}

// FakeEventBus – простая реализация для запоминания обработчиков и ручной публикации
type FakeEventBus struct {
	handlers map[string]func(any)
}

func NewFakeEventBus() *FakeEventBus {
	return &FakeEventBus{
		handlers: make(map[string]func(any)),
	}
}

func (f *FakeEventBus) Subscribe(event string, handler func(any)) {
	f.handlers[event] = handler
}

func (f *FakeEventBus) Publish(event string, data any) {
	if h, ok := f.handlers[event]; ok {
		h(data)
	}
}

func (f *FakeEventBus) GetHandler(event string) func(any) {
	return f.handlers[event]
}
