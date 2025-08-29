package wordle

import (
	"math"
)

type number interface {
	int | float32 | float64
}

func sum[T number](arr []T) T {
	var ret T
	for i := range arr {
		ret += arr[i]
	}
	return ret
}

func sumColumn[T number](arr [][]T, colNum int) T {
	var ret T
	for _, row := range arr {
		ret += row[colNum]
	}
	return ret
}

func DecToTer(i int) []int {
	var res []int
	iter := 0
	for i > 0 {
		v := i % 3
		res = append(res, v)
		i = i - v
		i = i / 3
		iter += 1
	}
	for len(res) < 5 {
		res = append(res, 0)
	}
	return res
}

func letterProbs(words [][]rune) [][]float64 {
	numWords := float64(len(words))
	var totalLetters float64 = float64(numWords) * 5

	counts := make([][]float64, 26)
	for i := range counts {
		counts[i] = make([]float64, 5)
	}
	letterOccur := make(map[rune]int)

	for _, word := range words {
		for j, letter := range word {
			counts[letter-97][j] += 1
			letterOccur[letter] += 1
		}
	}

	probs := make([][]float64, 26)
	for i := range probs {
		probs[i] = make([]float64, 5)
	}

	for i, row := range counts {
		for j := range row {
			placeGivenLetter := counts[i][j] / float64(letterOccur[rune(i+97)])
			probLetter := float64(letterOccur[rune(i+97)]) / totalLetters
			probs[i][j] = placeGivenLetter * probLetter / 0.2
		}
	}
	return probs
}

func findWordEntropy(probs [][]float64, word []rune) float64 {
	var prob float64
	for i, letter := range word {
		p := probs[letter-97][i]
		prob += p * math.Log2(p)
	}
	return -1 * prob
}

func wordProb(probs [][]float64, word []rune) float64 {
	var prob float64
	for i, letter := range word {
		prob += probs[letter-97][i]
	}
	return prob
}

func findCollectionEntropy(words [][]rune) float64 {
	probs := letterProbs(words)
	numWords := len(words)
	if numWords == 0 {
		return float64(0)
	}
	var H float64
	for _, word := range words {
		p := wordProb(probs, word) + 0.000001
		H += p * math.Log2(p)
	}
	return -1 * H
}

func FindBestWord(words [][]rune) ([]rune, float64) {
	// the number of ways the game could respond to us
	// TODO: this is the upper bound on the  number of ways
	// the game could respond, so we're doing more work
	// than we need to ususally
	// we could for example eliminate but letters we've locked in
	numResponses := int(math.Pow(3, 5))
	var bestWord []rune
	var bestEntropy float64 = math.MaxFloat64
	for _, word := range words {
		var entropy float64
		for i := 0; i < numResponses-1; i++ {
			response := DecToTer(i)
			newWords := FilterWords(words, word, response)
			entropy += findCollectionEntropy(newWords)
		}
		if entropy < bestEntropy {
			bestEntropy = entropy
			bestWord = word
		}
	}
	return bestWord, bestEntropy
}
