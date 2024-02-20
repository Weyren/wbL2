package main

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

import (
	"fmt"
)

type shape interface {
	printType()
	accept(visitor)
}

type square struct {
	//
}

func (s *square) printType() {
	fmt.Println("square")
}

func (s *square) accept(v visitor) {
	v.acceptForSquare(s)
}

type circle struct {
	//
}

func (c *circle) printType() {
	fmt.Println("circle")
}

func (c *circle) accept(v visitor) {
	v.acceptForCircle(c)
}

type visitor interface {
	acceptForSquare(*square)
	acceptForCircle(*circle)
}

type areaCalculator struct {
	//
}

func (a *areaCalculator) acceptForSquare(s *square) {
	fmt.Println("area square")
}

func (a *areaCalculator) acceptForCircle(c *circle) {
	fmt.Println("area circle")
}

func main() {
	square := &square{}
	circle := &circle{}
	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
}
