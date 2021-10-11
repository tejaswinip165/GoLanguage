package main

import (
	"fmt"
)

func sum(array []int) int {
	result := 0
	for v := range array {
		result += v
	}
	return result
}
func main() {
	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your array is: ", arr)
	fmt.Print("Result :", sum(arr))
}
