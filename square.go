package main

import "fmt"

func main() {
	var num, square int

	fmt.Println("Enter number : ")
	fmt.Scanln(&num)
	square = num * num
	fmt.Println("Square of  ", num, "is = ", square)
}
