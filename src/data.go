package wordle

import (
	"bufio"
	"os"
)

func filter[T any](array []T, testFunc func(T) bool) []T {
	var ret []T
	for _, item := range array {
		if testFunc(item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func listContains[T comparable](list []T, target T) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func LoadWords(filename string) ([][]rune, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var words [][]rune
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		words = append(words, []rune(scanner.Text()))
	}
	return words, scanner.Err()
}

func removeWordsWithChar(words [][]rune, char rune) [][]rune {
	test := func(s []rune) bool { return !listContains(s, char) }
	return filter(words, test)
}

func removeWordsWithPlacedChar(words [][]rune, char rune, index int) [][]rune {
	test := func(s []rune) bool { return !(s[index] == char) }
	return filter(words, test)
}

func removeWordsWithoutChar(words [][]rune, char rune) [][]rune {
	test := func(s []rune) bool { return listContains(s, char) }
	return filter(words, test)
}

func removeWordsWithoutPlacedChar(words [][]rune, char rune, index int) [][]rune {
	test := func(s []rune) bool { return (s[index] == char) }
	return filter(words, test)
}

func FilterWords(words [][]rune, guess []rune, response []int) [][]rune {
	for i, item := range response {
		switch item {
		case 0:
			words = removeWordsWithChar(words, guess[i])
		case 1:
			words = removeWordsWithoutChar(words, guess[i])
			words = removeWordsWithPlacedChar(words, guess[i], i)
		case 2:
			words = removeWordsWithoutPlacedChar(words, guess[i], i)
		}
	}
	return words
}
