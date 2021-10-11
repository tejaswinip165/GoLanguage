package main

import "fmt"

type item struct {
	itemid    int
	itemName  string
	itemPrice float64
	itemQ     int
}

func main() {
	var i item = item{}
	fmt.Println("itemid = ", i.itemid)
	fmt.Println("item name = ", i.itemName)
	fmt.Println("itemPrice = ", i.itemPrice)
	fmt.Println("Quntity = ", i.itemQ)

	i.itemid = 1000
	i.itemName = "ABC"
	i.itemPrice = 600
	i.itemQ = 20

	fmt.Println("Item ", i)

	var i2 item = item{itemName: "XYZ", itemPrice: 20, itemid: 200, itemQ: 200}
	fmt.Println(i2)

}
