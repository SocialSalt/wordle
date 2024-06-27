package main

import (
	"fmt"
	wordle "github.com/socialsalt/wordle/src"
	"strconv"
	"strings"
)

func main() {
	words, err := wordle.LoadWords("data/official_wordle_common.txt")
	if err != nil {
		panic("failed to load words")
	}
	bestWord, bestProb := wordle.FindBestWord(words)
	fmt.Println(string(bestWord))
	fmt.Println(bestProb)
	done := false
	for !done {
		var g string
		var r string
		var guess []rune
		var result []int
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&g)
		fmt.Print("Enter the result: ")
		fmt.Scanln(&r)
		guess = []rune(g)
		for _, token := range strings.Split(r, "") {
			number, _ := strconv.Atoi(token)
			result = append(result, number)
		}
		if r == "22222" {
			fmt.Println("You did it!")
			done = true
			continue
		}
		words = wordle.FilterWords(words, guess, result)
		bestWord, bestProb := wordle.FindBestWord(words)
		fmt.Println(string(bestWord))
		fmt.Println(bestProb)
		if len(words) < 50 {
			fmt.Println(words)
		}
		if bestProb == 0 {
			fmt.Println("This is our only guess, good luck!")
		}
	}
}
