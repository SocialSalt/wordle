package wordle

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

const TARGET_WORD_NULL_RUNE = '0'
const GUESS_WORD_NULL_RUNE = '1'

func assertPanic(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func assertError(condition bool, message string) error {
	if !condition {
		return errors.New(message)
	}
	return nil
}

func findFirst[T comparable](list []T, item T) int {
	for i, listItem := range list {
		if listItem == item {
			return i
		}
	}
	return -1
}

func checkWord(targetWord []rune, guess []rune) []string {
	res := make([]string, 5)
	for i := range res {
		res[i] = "🟥"
	}
	for i, letter := range guess {
		if targetWord[i] == letter {
			targetWord[i] = TARGET_WORD_NULL_RUNE
			guess[i] = GUESS_WORD_NULL_RUNE
			res[i] = "🟩"
		}
	}
	for i, letter := range guess {
		if index := findFirst(targetWord, letter); index != -1 {
			targetWord[index] = TARGET_WORD_NULL_RUNE
			res[i] = "🟨"
		}
	}
	return res
}

func PlayWordle(targetWord string, allowedWords []string) {
	gameOver := false
	for !gameOver {
		var guess string
		fmt.Println("Enter your guess")
		fmt.Scanln(&guess)
		if err := assertError(len(guess) == 5, "Guess must be 5 characters"); err != nil {
			fmt.Println(err)
			fmt.Println(len(guess))
			continue
		}
		if !slices.Contains(allowedWords, guess) {
			fmt.Printf("%s not found in dictionary\n", guess)
			continue
		}
		result := checkWord([]rune(targetWord), []rune(guess))
		fmt.Println(strings.Join(result, ""))
		gameOver = true
		for _, res := range result {
			if res != "🟩" {
				gameOver = false
			}
		}
	}
}
