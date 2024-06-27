package main

import (
	"fmt"
	wordle "github.com/socialsalt/wordle/src"
	"math/rand"
)

func main() {
	words, err := wordle.LoadWords("words.txt")
	if err != nil {
		panic(err)
	}
	index := rand.Intn(len(words))
	targetWord := string(words[index])
	fmt.Printf("You found the target world %s\n", targetWord)
	wordle.PlayWordle(targetWord)
	fmt.Printf("You found the target world %s\n", targetWord)
}
