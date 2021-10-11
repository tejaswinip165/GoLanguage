package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "jay jawan jay kisan"
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		fmt.Println(char)
	}
	res1 := strings.Count(str, "j")
	fmt.Println(" J = ", res1)

	res2 := strings.Count(str, "a")
	fmt.Println(" A = ", res2)

	res3 := strings.Count(str, "s")
	fmt.Println(" S = ", res3)
}
