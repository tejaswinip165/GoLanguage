package main

import "fmt"

func main() {
	var alpha string
	fmt.Println("enter alphabate")
	fmt.Scanln(&alpha)
	switch alpha {
	case "a":
		fmt.Println("Its vowel")
	case "e":
		fmt.Println("Its vowel")
	case "i":
		fmt.Println("Its vowel")
	case "o":
		fmt.Println("Its vowel")
	case "u":
		fmt.Println("Its vowel")
	default:
		fmt.Println("Its Alphabate")
	}
}
