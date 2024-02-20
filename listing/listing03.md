Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Вывод: nil false
fmt.Println(err) печатает значение (value) интерфейса err - nil так как не было задано 
fmt.Println(err == nil) сравнивет интерфейс с nil - false так как есть информация о типе PathError
который имеет метод Error() string и реализует интерфейс error

Интерфейс хранит указатель на данные(значение интерфейса) и таблицу с информацией об объекте(методы, тип)
type iface struct {
	tab  *itab
	data unsafe.Pointer
}
```
