package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50}
	fmt.Println(x)

	last := x[len(x)-1]

	fmt.Printf("Last element: %v\n", last)

}
