package main

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

import "fmt"

type request string

type iHandler interface {
	process(*request)
	setNextHandle(iHandler)
}

type processRequest struct {
	handler iHandler
}

func (po *processRequest) process(r *request) {
	po.handler.process(r)
}

type concreteHandle1 struct {
	request
	next iHandler
}

func (ch1 *concreteHandle1) process(r *request) {
	if *r == ch1.request {
		fmt.Println("process first handler")
		return
	} else {
		ch1.next.process(r)
	}
}

func (ch1 *concreteHandle1) setNextHandle(h iHandler) {
	ch1.next = h
}

type concreteHandle2 struct {
	request
	next iHandler
}

func (ch2 *concreteHandle2) process(r *request) {
	if *r == ch2.request {
		fmt.Println("process second handler")
		return
	} else {
		ch2.next.process(r)
	}
}

func (ch2 *concreteHandle2) setNextHandle(h iHandler) {
	ch2.next = h
}

type concreteHandle3 struct {
	request
	next iHandler
}

func (ch3 *concreteHandle3) process(r *request) {
	if *r == ch3.request {
		fmt.Println("process three handler")
		return
	} else {
		ch3.next.process(r)
	}
}

func (ch3 *concreteHandle3) setNextHandle(h iHandler) {
	ch3.next = h
}

func main() {

	ch1 := &concreteHandle1{request: "first"}
	ch2 := &concreteHandle2{request: "second"}
	ch3 := &concreteHandle3{request: "three"}

	ch1.setNextHandle(ch2)
	ch2.setNextHandle(ch3)

	processOrder := &processRequest{handler: ch1}

	req := new(request)
	*req = "second"

	processOrder.process(req)
}
