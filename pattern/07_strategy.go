package main

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

import "fmt"

type strategy interface {
	execute(data string)
}

type strategyA struct{}

func (s *strategyA) execute(data string) {
	fmt.Println("execute strategy A")
}

type strategyB struct{}

func (s *strategyB) execute(data string) {
	fmt.Println("execute strategy B")
}

type context struct {
	strategy
}

func (c *context) setStrategy(s strategy) {
	c.strategy = s
}

func (c *context) execute(data string) {
	c.strategy.execute(data)
}

func main() {

	strategyA := &strategyA{}
	strategyB := &strategyB{}

	context := &context{}

	context.setStrategy(strategyA)
	context.execute("strategy")

	context.setStrategy(strategyB)
	context.execute("strategy")
}
