package main

import "fmt"

func main() {
	var size, i int

	fmt.Print("Enter the Even Odd Array Size = ")
	fmt.Scan(&size)

	evoddarr := make([]int, size)

	fmt.Print("Enter the Even Odd Array Items  = ")
	for i = 0; i < size; i++ {
		fmt.Scan(&evoddarr[i])
	}

	evenCount := 0
	oddCount := 0

	for i = 0; i < size; i++ {
		if evoddarr[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	fmt.Println("The Total Number of Even Numbers = ", evenCount)
	fmt.Println("The Total Number of Odd Numbers  = ", oddCount)
}
