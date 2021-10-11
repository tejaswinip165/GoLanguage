// Golang program to count the number of
// repeating words in given Golang String
package main

import (
	"fmt"
	"strings"
)

func repetition(st string) map[string]int {

	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

func main() {
	input := "jay jawan jay kisan "

	for index, element := range repetition(input) {
		fmt.Println(index, "=", element)
	}
}
