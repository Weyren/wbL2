package main

import "testing"

func TestSearchAnagrams(t *testing.T) {
	testTable := []struct {
		input       *[]string
		exp         *map[string]*[]string
		description string
	}{
		{
			input:       &[]string{"пятак"},
			exp:         &map[string]*[]string{},
			description: "словарь - 1 слово",
		},
		{
			input:       &[]string{"пятка", "пятак"},
			exp:         &map[string]*[]string{"пятка": &[]string{"пятак", "пятка"}},
			description: "сортировка множества",
		},
		{
			input:       &[]string{"лось", "пятка", "пятак", "соль"},
			exp:         &map[string]*[]string{"пятка": {"пятак", "пятка"}, "лось": {"лось"}},
			description: "сортировка множества",
		},
	}

	for _, testCase := range testTable {
		result := SearchAnagrams(testCase.input)
		for k, _ := range *result {
			if resultSlice, exist1 := (*result)[k]; exist1 {
				if expSlice, exist2 := (*testCase.exp)[k]; exist2 {
					for i := 0; i < len(*resultSlice); i++ {
						if (*resultSlice)[i] != (*expSlice)[i] {
							t.Error("Неверный порядок")
						}
					}
				}
			}
		}
	}
}
