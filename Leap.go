package main

import "fmt"

func main() {
	var year int
	fmt.Println("Enter number : ")
	fmt.Scanln(&year)

	if year%400 == 0 {
		fmt.Println(year, " is a Leap Year")
	} else if year%100 == 0 {
		fmt.Println(year, " is Not a Leap Year")
	} else if year%4 == 0 {
		fmt.Println(year, " is a Leap Year")
	} else {
		fmt.Println(year, " is Not a Leap Year")
	}
}
