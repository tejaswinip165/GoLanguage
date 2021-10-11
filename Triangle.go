package main

import "fmt"

func main() {
	var base, height, area float64
	fmt.Println("Enter base and height : ")
	fmt.Scanln(&base, &height)
	area = (base * height) / 2

	fmt.Println("Area = ", area)

}
