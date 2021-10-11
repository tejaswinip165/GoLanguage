package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println("hello")
	var num1, num2 int
	fmt.Print("enter num1 : ")
	fmt.Scanln(&num1)
	fmt.Println("Enter num 2: ")
	fmt.Scanln(&num2)
	fmt.Println(num1 + num2)

}
