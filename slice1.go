package main

import "fmt"

func main() {

	// create slice using make method

	s := make([]int, 5, 5)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[3] = 40
	s[4] = 50
	fmt.Println("Slice : ", s)

	for i := 0; i < len(s); i++ {
		fmt.Println(i, "=", s[i])
	}

}
