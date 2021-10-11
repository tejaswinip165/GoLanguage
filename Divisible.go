package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter number : ")
	fmt.Scanln(&num)
	if num%5 == 0 && num%11 == 0 {
		fmt.Println("Number D by 5 and 11")
	} else {
		fmt.Println("Number is not D by 5 and 11")
	}
}
