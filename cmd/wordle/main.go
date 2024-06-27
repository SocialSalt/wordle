package main

import (
	"fmt"
	"math/rand"
	"wordle/src"
)

func main() {
	words, err := wordle.LoadWords("words.txt")
	if err != nil {
		panic(err)
	}
	index := rand.Intn(len(words))
	targetWord := words[index]
	fmt.Printf("You found the target world %s\n", targetWord)
	wordle.PlayWordle(targetWord)
	fmt.Printf("You found the target world %s\n", targetWord)
}
