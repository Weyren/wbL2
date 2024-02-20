package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern


*/

//есть несколько фоматов документа,
//фабрика создает объект в соответствии с типом

type Json struct {
	data   []byte
	kind   string
	lenght int
}

func (j Json) GetInfo() {
	fmt.Println(j.kind, j.lenght)
}

type Yaml struct {
	data    []byte
	kind    string
	version string
}

func (y Yaml) GetInfo() {
	fmt.Println(y.kind, y.version)
}

type Document interface {
	GetInfo()
}

func NewDocument(kind string) Document {
	switch kind {
	case "json":
		return Json{
			[]byte{},
			"json",
			0,
		}
	case "yaml":
		return Yaml{
			data:    []byte{},
			kind:    "yaml",
			version: "1.0",
		}
	default:
		fmt.Println("incorrect kind")
		return nil
	}
}

func main() {
	documenst := []string{"json", "yaml", "xml"}
	for _, kind := range documenst {
		doc := NewDocument(kind)
		if doc == nil {
			continue
		}
		doc.GetInfo()
	}
}
