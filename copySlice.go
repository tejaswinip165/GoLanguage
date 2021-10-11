package main

import "fmt"

func main() {

	slice1 := []string{"ABC", "DEF", "GHI"}
	slice2 := make([]string, 3)
	// Before Copying
	fmt.Println("Befor copying Slice1 =  ", slice1)
	fmt.Println("Before copying Slice2: ", slice2)
	Copy1 := copy(slice2, slice1)
	fmt.Println("After copying Slice1 =  ", slice1)
	fmt.Println("after copying Slice2 = ", slice2)
	fmt.Println("Number of elements copied: ", Copy1)
}
