package main

import "fmt"

func main() {
	var days, months, week, year int
	fmt.Println("enter days : ")
	fmt.Scanln(&days)
	months = days / 30
	week = days / 7
	year = days / 365
	fmt.Println("days = ", days, "months=", months, "week= ", week, "year = ", year)

}
