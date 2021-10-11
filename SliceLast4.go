package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50} // Intialized with values
	fmt.Println(x)

	last := x[len(x)-1]
	last1 := x[len(x)-2]
	last2 := x[len(x)-3]
	last3 := x[len(x)-4]
	fmt.Printf("Last element: %v\n", last)
	fmt.Printf("2nd Last element: %v\n", last1)
	fmt.Printf("3rd Last element: %v\n", last2)
	fmt.Printf("4th Last element: %v\n", last3)
	fmt.Printf("Last 4 elements: %v,%v,%v,%v", last, last1, last2, last3)

}
