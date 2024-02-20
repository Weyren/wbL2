package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// SearchAnagrams возвращает мапу множеств анаграмм
func SearchAnagrams(elems *[]string) *map[string]*[]string {

	anagrams := make(map[string]*[]string) //мапа анаграм

	for _, elem := range *elems {
		elem = strings.ToLower(elem) //нижний регистр
		sortedElem := sortWord(elem) //сортировка строки

		if _, exist := anagrams[sortedElem]; !exist { //если ключа нет - создаем слайс из 1 элемента и кладем в значение ссылку
			anagrams[sortedElem] = &[]string{elem}
		} else { //добавление по ключу слова в слайс анаграмм
			*anagrams[sortedElem] = append(*anagrams[sortedElem], elem)
		}
	}

	result := make(map[string]*[]string)
	for _, words := range anagrams {
		correctKey := (*words)[0]           //первое в словаре слово для конкретного множества
		words = removeDuplicateWords(words) //удалили дубликаты

		if len(*words) > 1 { //анаграмм нет если всего 1 значение в слайсе - удаляем
			sort.Strings(*words)       // сортировка строк по алфавиту
			result[correctKey] = words //запись в результирующую мапу
		}

	}

	return &result
}

// sortWord сортирует символы в строке в алфавитном порядке
func sortWord(word string) string {
	rword := []rune(word)
	sort.Slice(rword, func(i, j int) bool {
		return rword[i] < rword[j]
	})

	return string(rword)
}

// removeDuplicateWords удаляет повторы
func removeDuplicateWords(words *[]string) *[]string {
	m := make(map[string]struct{}, 0) //мапа уникальных слов
	result := make([]string, 0)       //слайс уникальных слов

	for _, v := range *words {
		if _, exists := m[v]; !exists {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}
	return &result
}

func main() {
	a := *SearchAnagrams(&[]string{"пятак", "пятка", "тяпка", "соль", "лось"})
	for k, v := range a {
		fmt.Println(k, *v)
	}
}
