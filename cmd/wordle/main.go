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
	index := rand.Intn(len(words))
	targetWord := string(words[index])
	wordle.PlayWordle(targetWord)
	fmt.Printf("You found the target world %s\n", targetWord)
}
