package main

import "fmt"

func main() {
	// concat 2 slice
	var slice1 = []string{"ABC", "DEF", "GHI"}
	var slice2 = []string{"JKL", "MNO", "PQR", "STU"}
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	var concSlice = []string{}
	concSlice = append(slice1, slice2...)
	fmt.Println("Slice after Concatination = ", concSlice)
}
