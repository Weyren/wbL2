package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Collector interface {
	SetCore()
	SetMemory()
	SetBrand()
	GetComputer() Computer
}

func GetCollector(ct string) Collector {
	switch ct {
	case "asus":
		return &AsusCollector{}
	case "hp":
		return &HpCollector{}
	default:
		return nil
	}
}

type AsusCollector struct {
	Core   int
	Memory int
	Brand  string
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}
func (collector *AsusCollector) SetMemory() {
	collector.Memory = 8
}
func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:   collector.Core,
		Memory: collector.Memory,
		Brand:  collector.Brand,
	}
}

type HpCollector struct {
	Core   int
	Memory int
	Brand  string
}

func (collector *HpCollector) SetCore() {
	collector.Core = 8
}
func (collector *HpCollector) SetMemory() {
	collector.Memory = 16
}
func (collector *HpCollector) SetBrand() {
	collector.Brand = "HP"
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		Core:   collector.Core,
		Memory: collector.Memory,
		Brand:  collector.Brand,
	}
}

type Computer struct {
	Core   int
	Memory int
	Brand  string
}

func (pc *Computer) Print() {
	fmt.Println(pc.Core, pc.Memory, pc.Brand)
}

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

func (factory *Factory) CreateComputer() Computer {
	factory.Collector.SetMemory()
	factory.Collector.SetCore()
	factory.Collector.SetBrand()
	return factory.Collector.GetComputer()
}
func main() {
	asusCollector := GetCollector("asus")
	factory := NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	hpCollector := GetCollector("hp")
	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()

}
