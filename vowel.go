package main

import "fmt"

func main() {
	var alp string
	fmt.Println("Enter alphabate : ")
	fmt.Scanln(&alp)
	if alp == "a" || alp == "e" || alp == "i" || alp == "o" || alp == "u" {
		fmt.Println("Its vowel")
	} else {
		fmt.Println("its consonant")
	}
}
