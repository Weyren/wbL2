package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var (
	flagFilepath   = flag.String("filepath", "", "filepath")
	flagPattern    = flag.String("pattern", "", "substring to search")
	flagAfter      = flag.Int("A", 0, "печатать +N строк после совпадения")
	flagBefore     = flag.Int("B", 0, "печатать +N строк до совпадения")
	flagContext    = flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	flagCount      = flag.Bool("c", false, "печатать количество строк")
	flagIrnoreCase = flag.Bool("i", false, "игнорировать регистр")
	flagInvert     = flag.Bool("v", false, "вместо совпадения, исключать")
	flagFixed      = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	flagLineNum    = flag.Bool("n", false, "печатать номер строки")
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

	data := strings.Split(string(fileData), "\n")
	if *flagPattern == "" {
		return errors.New("no string to search")
	}

	if *flagAfter > 0 {
		for i, v := range data {
			if strings.Contains(v, *flagPattern) {
				if i+*flagAfter+1 < len(data) {
					fmt.Println(data[i : i+*flagAfter+1])
				} else {
					fmt.Println(errors.New("не хватает строк"))
				}
			}
		}
	}

	if *flagBefore > 0 {
		for i, v := range data {
			if strings.Contains(v, *flagPattern) {
				if i-*flagBefore >= 0 {
					fmt.Println(data[i-*flagBefore : i+1])
				} else {
					fmt.Println(errors.New("не хватает строк"))
				}
			}
		}
	}

	if *flagContext > 0 {
		for i, v := range data {
			if strings.Contains(v, *flagPattern) {
				if i+*flagContext+1 < len(data) && i-*flagContext >= 0 {
					fmt.Println(data[i-*flagContext : i+*flagContext+1])
				} else {
					fmt.Println(errors.New("не хватает строк"))
				}
			}
		}
	}

	if *flagCount == true {
		count := 0
		for _, v := range data {
			if strings.Contains(v, *flagPattern) {
				count++
			}
		}
		fmt.Println("Количество строк:", count)
	}

	if *flagIrnoreCase == true {
		for _, v := range data {
			if strings.Contains(strings.ToLower(v), strings.ToLower(*flagPattern)) {
				fmt.Println(v)
			}
		}
	}

	if *flagInvert == true {
		for i := 0; i < len(data)-1; i++ {
			if strings.Contains(data[i], *flagPattern) {
				data = append(data[:i], data[i+1:]...)
			}
		}
		fmt.Println("Строки удалены")
		fmt.Println(data)
	}

	if *flagFixed == true {
		for _, v := range data {
			if v == *flagPattern {
				fmt.Println(v)
			}
		}
	}

	if *flagLineNum == true {
		for i, v := range data {
			if strings.Contains(v, *flagPattern) {
				fmt.Println("Номер строки:", i+1)
			}
		}
	}
	return nil
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
