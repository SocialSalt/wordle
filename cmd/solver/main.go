package main

import (
	"fmt"
	"strconv"
	"strings"

	wordle "github.com/socialsalt/wordle/internal"
)

func main() {
	words, err := wordle.LoadWords("words/official_wordle_common.txt")
	if err != nil {
		panic("failed to load words")
	}
	fmt.Println("Preparing solver...")
	bestWord, bestProb := wordle.FindBestWord(words)
	fmt.Println(string(bestWord))
	fmt.Println(bestProb)
	for {
		var g string
		var r string
		var guess []rune
		var result []int
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&g)
		if len(g) != 5 {
			fmt.Println("guess must be of length 5")
			continue
		}
		fmt.Print("Enter the result: ")
		fmt.Scanln(&r)
		if len(r) != 5 {
			fmt.Println("result must be of length 5")
			continue
		}
		guess = []rune(g)
		for _, token := range strings.Split(r, "") {
			number, _ := strconv.Atoi(token)
			result = append(result, number)
		}
		if r == "22222" {
			fmt.Println("You did it!")
			return
		}
		words = wordle.FilterWords(words, guess, result)
		bestWord, bestProb := wordle.FindBestWord(words)
		fmt.Println(string(bestWord))
		fmt.Printf("%.3f\n", bestProb)
		if len(words) < 50 {
			for _, word := range words {
				fmt.Printf("%s ", string(word))
			}
			fmt.Printf("\n")
		}
		if bestProb == 0 {
			fmt.Println("This is our only guess, good luck!")
			return
		}
	}
}
