package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter number : ")
	fmt.Scanln(&num)
	if num == 0 {
		fmt.Println("The Number is Zero")
	}
	if num > 0 {
		fmt.Println("The Number is positive")
	}
	if num < 0 {
		fmt.Println("The Number is negative")
	}

}
