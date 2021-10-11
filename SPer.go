package main

import "fmt"

func main() {
	var sub1, sub2, sub3, sub4, sub5 int
	fmt.Println("Enter marks of the five subjects:")
	fmt.Scanf("%d", &sub1)
	fmt.Scanf("%d", &sub2)
	fmt.Scanf("%d", &sub3)
	fmt.Scanf("%d", &sub4)
	fmt.Scanf("%d", &sub5)
	avg := (sub1 + sub2 + sub3 + sub4 + sub5) / 5
	if avg >= 90 {
		print("Grade: A")
	} else if avg >= 80 && avg < 90 {
		print("Grade: B")
	} else if avg >= 70 && avg < 80 {
		print("Grade: C")
	} else if avg >= 60 && avg < 70 {
		print("Grade: D")
	} else {
		print("Grade: F")
	}
}
