package main

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import "fmt"

type iCommand interface {
	execute()
}

type concreteCommand struct {
	reciver *reciver
}

func (cc *concreteCommand) execute() {
	cc.reciver.action()
}

type reciver struct{}

func (r *reciver) action() {
	fmt.Println("execute concrete command")
}

type invoker struct {
	command iCommand
}

func (i *invoker) executeCommand() {
	i.command.execute()
}

func main() {
	reciver := &reciver{}
	command := &concreteCommand{reciver: reciver}
	invoker := &invoker{command: command}

	invoker.executeCommand()
}
