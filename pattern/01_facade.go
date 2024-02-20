package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// фасад карт
type cardFacade struct {
	cc creditCard
	dc debtCard
}

// кредитная
type creditCard struct {
	id      string
	balance int
}

// метод получения баланса
func (c creditCard) getDebt() {
	fmt.Println("debt", c.balance)
}

// дебетовая
type debtCard struct {
	id   string
	debt int
}

// метод получения долга
func (c debtCard) getBalance() {
	fmt.Println("balance", c.debt)
}

func main() {
	cc := creditCard{"1", 1000}         //структура кредитки
	dc := debtCard{"2", 200}            //структура дебетовой карты
	cards := cardFacade{cc: cc, dc: dc} //фасад для работы с картами
	cards.cc.getDebt()
	cards.dc.getBalance()
}
