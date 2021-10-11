package main

import "fmt"

func main() {
	var len, bre, area, perimeter float64
	fmt.Println("Enter l and b : ")
	fmt.Scanln(&len, &bre)
	area = len * bre
	perimeter = 2 * (len + bre)
	fmt.Println("Area = ", area)
	fmt.Println("Perimeter = ", perimeter)

}
