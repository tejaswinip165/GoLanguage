package main

import "fmt"

func main() {
	var radius, area, circmfrnc, diameter float64
	const pi float64 = 3.14
	fmt.Println("Enter radius of circle : ")
	fmt.Scanln(&radius)
	area = pi * radius * radius
	circmfrnc = 2 * radius * pi
	diameter = 2 * radius
	fmt.Println("area of circle = ", area)
	fmt.Println("Ci of circle = ", circmfrnc)
	fmt.Println("Diameter of circle =", diameter)

}
