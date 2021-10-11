package main

import "fmt"

func main() {
	var num1, num2, add, sub, div, mul, choice int
	fmt.Println("1.add")
	fmt.Println("2.subb")
	fmt.Println("3.div")
	fmt.Println("4.mul")
	fmt.Println("Enter Choice")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Num1 num2 = ")
		fmt.Scanln(&num1, &num2)
		add = num1 + num2
		fmt.Println("Addition =", add)
	case 2:
		fmt.Println("Num1 num2 = ")
		fmt.Scanln(&num1, &num2)
		sub = num1 - num2
		fmt.Println("Sub =", sub)
	case 3:
		fmt.Println("Num1 num2 = ")
		fmt.Scanln(&num1, &num2)
		div = num1 / num2
		fmt.Println("Div =", div)
	case 4:
		fmt.Println("Num1 num2 = ")
		fmt.Scanln(&num1, &num2)
		mul = num1 * num2
		fmt.Println("mul =", mul)

	}
}
