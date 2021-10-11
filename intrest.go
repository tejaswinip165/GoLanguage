package main

import "fmt"

func main() {
	var P, T, R, SI int

	fmt.Println("Enter P : ")
	fmt.Scanln(&P)
	fmt.Println("Enter T: ")
	fmt.Scanln(&T)
	fmt.Println("Enter R : ")
	fmt.Scanln(&R)
	SI = P * R * T
	fmt.Println("Simple intresr =", SI)
}
