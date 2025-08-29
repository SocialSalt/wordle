package wordle

import (
	"errors"
	"fmt"
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

func checkWord(targetWord []rune, guess []rune) []int {
	res := make([]int, 5)
	for i := range res {
		res[i] = 0
	}
	for i, letter := range guess {
		if targetWord[i] == letter {
			targetWord[i] = TARGET_WORD_NULL_RUNE
			guess[i] = GUESS_WORD_NULL_RUNE
			res[i] = 2
		}
	}
	for i, letter := range guess {
		if index := findFirst(targetWord, letter); index != -1 {
			targetWord[index] = TARGET_WORD_NULL_RUNE
			res[i] = 1
		}
	}
	return res
}

func PlayWordle(targetWord string) {
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
		result := checkWord([]rune(targetWord), []rune(guess))
		fmt.Println(result)
		gameOver = true
		for _, res := range result {
			if res != 2 {
				gameOver = false
			}
		}
	}
}
