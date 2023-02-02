package main

import (
	"unicode"
	"fmt"
)

func IsIsogram(word string) bool {
	// create a map to store chars and no of occurrences

	letterMap := make(map[string]int)

	// loop through word, store letter and increase occurrence

	for _, value := range word {
		letter := string(value)
		if letterMap[letter] > 0 {
			letterMap[letter] += 1
		} else {
			letterMap[letter] = 1
		}
	}
	fmt.Println(letterMap)

	// loop through object, anything over 2?
	for item := range letterMap {
		if letterMap[item] > 1 && unicode.IsLetter(item) {
			fmt.Printf(`this is a letter occuring more than once - %v`, item)
		}
	}

	// if over 2 and is not a letter (ascii value? in built function?)
	// true
	// anything over 2 is a letter - false
	return false

}

func main() {
	fmt.Println(IsIsogram("backgrobbbund"))
}
