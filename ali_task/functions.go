package main

// this file includes all functions that perform the required logic
import (
	"log"
	"os"
	"strings"
	"unicode"
)

func readFile(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error in reading file")
	}
	return string(data), nil
}
func totalWordsCount(fileContent string) int {
	words := strings.FieldsFunc(fileContent, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	return len(words)
}

func MostOccurredWord(fileContent string) []string {
	totalWords := make(map[string]int)
	words := strings.FieldsFunc(fileContent, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	for _, words := range words {
		totalWords[strings.ToLower(words)]++
	}

	maxCount := 0
	var mostOccurred []string
	for word, count := range totalWords {
		if count > maxCount {
			maxCount = count
			mostOccurred = []string{word}
		} else if count == maxCount {
			mostOccurred = append(mostOccurred, word)
		}
	}

	return mostOccurred

}

func fileCharacterCount(fileContent string) int {
	return len(fileContent)
}

func averageSentenceLength(fileContent string) float64 {

	sentences := strings.FieldsFunc(fileContent, func(r rune) bool {
		return r == '.' || r == '!' || r == '?'
	})
	words := strings.FieldsFunc(fileContent, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	totalWords := len(words)
	totalSen := len(sentences)

	avgLength := 0.0
	if totalSen > 0 {
		avgLength = float64(totalWords) / float64(totalSen)
	}

	return avgLength
}
