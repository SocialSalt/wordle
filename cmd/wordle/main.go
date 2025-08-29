package main

import (
	"fmt"
	wordle "github.com/socialsalt/wordle/internal"
	"math/rand"
)

func main() {
	words, err := wordle.LoadWords("words/official_wordle_common.txt")
	if err != nil {
		panic(err)
	}
	allowedWords, err := wordle.LoadWordsString("words/words.txt")
	index := rand.Intn(len(words))
	targetWord := string(words[index])
	wordle.PlayWordle(targetWord, allowedWords)
	fmt.Printf("You found the target word %s\n", targetWord)
}
