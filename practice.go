package main

import (
	"fmt"
	"strings"
)

func main() {

	test1 := "catcat"
	test2 := "dog"
	test3 := "catcat"

	//strings are equal or not
	fmt.Println("test1:", test1)
	fmt.Println("test2:", test2)

	if test1 == test2 {
		fmt.Println("Strings are equal")
	}

	if test1 != test2 {
		fmt.Println("Strings are not equal")
	}

	//check using strings function
	fmt.Println(strings.EqualFold(test1, test2))
	fmt.Println(strings.EqualFold(test3, test1))

	// lenght of string
	var length = len([]rune(test1))
	println("Length of the string is : ", length)

	// title case
	fmt.Println(strings.ToTitle("tejaswini vasant pawar"))

	//upper case
	fmt.Println(strings.ToUpper(test2))

	// lower case
	fmt.Println(strings.ToLower(test2))

	//get substring
	myString := "india is my country;)"

	a := []rune(myString)

	myShortString := string(a[0:6])

	fmt.Println(myShortString)

	//replace string by other string

	fmt.Println(strings.Replace("india is the best ", "best", "best country", 2))

	//concate strings
	result := fmt.Sprintf("%s%s%s", test1,
		test2, test3)

	fmt.Println(result)

}
