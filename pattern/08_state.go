package main

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import "fmt"

// Интерфейс, представляющий состояние
type State interface {
	Handle() State
}

// Структура, представляющая контекст
type Context struct {
	state State
}

// Метод для установки состояния
func (c *Context) setState(state State) {
	c.state = state
}

// Выполнение действия
func (c *Context) request() {
	c.setState(c.state.Handle())
}

// Реализация первого состояния
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() State {
	fmt.Println("Выполняется действие в состоянии A")
	return &ConcreteStateB{}
}

// Реализация второго состояния
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() State {
	fmt.Println("Выполняется действие в состоянии B")
	return &ConcreteStateA{}
}

func main() {
	// Создание контекста
	context := &Context{}

	// Установка начального состояния
	initialState := &ConcreteStateA{}
	context.setState(initialState)

	// Выполнение действия
	context.request() // Вывод: "Выполняется действие в состоянии A"

	// Изменение состояния и выполнение действия снова
	context.request() // Вывод: "Выполняется действие в состоянии B"
}
