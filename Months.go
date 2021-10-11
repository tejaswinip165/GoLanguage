package main

import "fmt"

func main() {
	var month int
	fmt.Println("1.Jan 2.feb 3.Mar 4.Apr 5.may 6.june 7.July 8.Aug 9.Sep 10.Oct 11.Nov 11.dec")
	fmt.Println("Enter No of month")
	fmt.Scanln(&month)
	switch month {
	case 1:
		fmt.Println(" Jan= 31 Days")
	case 2:
		fmt.Println("Feb = 28 Days")
	case 3:
		fmt.Println("March = 31 Days")
	case 4:
		fmt.Println("April = 30 Days")
	case 5:
		fmt.Println("may = 31 Days")
	case 6:
		fmt.Println("Jun = 30 Days")
	case 7:
		fmt.Println(" july = 31 Days")
	case 8:
		fmt.Println("31 Days")
	case 9:
		fmt.Println("30 Days")
	case 10:
		fmt.Println("31 Days")
	case 11:
		fmt.Println("30 Days")
	case 12:
		fmt.Println("31 Days")

	}
}
