package main

import "fmt"

func main() {
	var cities = []string{"Mumbai", "Pune", "nagapur", "Satara"}
	fmt.Println(cities)

	cities = RemoveIndex(cities, 3)
	fmt.Println(cities)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
