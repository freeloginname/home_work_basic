package countwords

import (
	"fmt"
	"strings"
)

const (
	punctuation = ",.!?;"
)

func countWords(inputString string) map[string]int {
	var wordsAppearance = make(map[string]int)
	for _, word := range strings.Split(inputString, " ") {
		if strings.ContainsAny(word, punctuation) {
			var fixWord string
			var wordRunes = []rune(word)
			for index := len(wordRunes) - 1; index >= 0; index-- {
				if !strings.ContainsAny(string(wordRunes[index]), punctuation) {
					fixWord = word[0 : index+1]
					break
				}
			}
			if fixWord != "" && len(fixWord) > 0 {
				fmt.Println(fixWord)
				wordsAppearance[fixWord]++
			}
		} else if word != "" && len(word) > 0 {
			fmt.Println(word)
			wordsAppearance[word]++
		}
	}
	return wordsAppearance
}

func PublicCountWords(inputString string) map[string]int {
	return countWords(inputString)
}
