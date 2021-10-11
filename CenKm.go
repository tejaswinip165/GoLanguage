package main

import "fmt"

func main() {
	var cen, met, km float64
	fmt.Println("enter cent : ")
	fmt.Scanln(&cen)
	met = cen / 1000
	km = cen / 100000
	fmt.Println("meter = ", met, "Km =", km)

}
