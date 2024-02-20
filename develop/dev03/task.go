package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки +
-n — сортировать по числовому значению +
-r — сортировать в обратном порядке +
-u — не выводить повторяющиеся строки +

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные +
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	flagFilepath = flag.String("filepath", "", "filepath")
	flagColumn   = flag.Int("k", 0, "set column to sort")
	flagIntValue = flag.Bool("n", false, "sort by int value")
	flagReversed = flag.Bool("r", false, "reversed sort")
	flagUnic     = flag.Bool("u", false, "no repeat sort")
	flagChek     = flag.Bool("c", false, "check if data sorted")
)

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}
	flag.Parse()

	fileData, err := os.ReadFile(*flagFilepath)
	if err != nil {
		return err
	}

	var data [][]string

	fileStrings := strings.Split(string(fileData), "\n")
	for _, val := range fileStrings {
		data = append(data, strings.Fields(val))
	}

	//-k
	if *flagColumn > 0 {
		*flagColumn--
		sort.Slice(data, func(i, j int) bool {
			if *flagReversed == true {
				return !(data[i][*flagColumn] < data[j][*flagColumn])
			}
			return data[i][*flagColumn] < data[j][*flagColumn]
		})
		*flagColumn = 0
		fmt.Println(data)
	}

	//-n
	if *flagIntValue == true {
		sort.Slice(data, func(i, j int) bool {
			levo, err := strconv.Atoi(data[i][*flagColumn])
			if err != nil {
				panic(err)
			}
			pravo, err := strconv.Atoi(data[j][*flagColumn])
			if err != nil {
				panic(err)
			}

			if *flagReversed == true {
				return !(levo > pravo)
			}

			return levo < pravo

		})
		*flagIntValue = false
		fmt.Println(data)
	}

	if *flagUnic == true {
		sort.Slice(data, func(i, j int) bool {
			if *flagReversed == true {
				return !(data[i][*flagColumn] < data[j][*flagColumn])
			}
			return data[i][*flagColumn] < data[j][*flagColumn]
		})

		unicStrings := make(map[string]struct{})
		result := make([][]string, 0)

		for _, val := range data {
			valstr := strings.Join(val, " ")
			_, exists := unicStrings[valstr]

			if exists == false {
				result = append(result, val)
				unicStrings[valstr] = struct{}{}
			}
		}
		*flagUnic = false

		fmt.Println(result)
	}

	if *flagChek == true {
		sorted := sort.SliceIsSorted(fileStrings, func(i, j int) bool {
			return fileStrings[i][*flagColumn] < fileStrings[j][*flagColumn]
		})
		fmt.Println("Sorted:", sorted)
		*flagChek = false
	}
	return nil
}

func main() {

	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
