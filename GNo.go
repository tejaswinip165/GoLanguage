package main

import "fmt"

func main() {
	var no1, no2, no3 int
	fmt.Println("Enter 3 numbers : ")
	fmt.Scanln(&no1, &no2, &no3)
	if no1 >= no2 && no1 >= no3 {
		fmt.Println("Number 1 ", no1, "is greatest")
	}
	if no2 >= no3 && no2 >= no1 {
		fmt.Println("Number 2 ", no2, "is greatest")
	}
	if no3 >= no2 && no3 >= no1 {
		fmt.Println("Number 3 ", no3, "is greatest")
	}

}
