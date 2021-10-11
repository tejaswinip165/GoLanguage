package main

import (
	"fmt"
	f "fmt"
	"strings"
)

func main() {
	var str1 = "mahendra singh dhoni"
	f.Println("String 1 = ", str1)
	var str2 = "Sachin ramesh tendulkar"
	f.Println("String 2 = ", str2)

	f.Println(strings.ContainsAny(str1, str2))
	f.Println(strings.ContainsAny(str1, "singh"))
	f.Println(strings.ContainsAny(str2, "T"))

	fmt.Println(strings.EqualFold(str1, str2))
	fmt.Println(strings.EqualFold("mahendra singh dhoni", "Mahendra singh dhoni"))

	array := strings.Fields(str1)
	for _, v := range array {
		fmt.Println(v)
	}

	fmt.Println(strings.HasPrefix(str1, "mah"))
	fmt.Println(strings.HasSuffix(str2, "kar"))
	fmt.Println(strings.Index(str1, "d"))
	fmt.Println(strings.Index(str2, "d"))

	fmt.Println(strings.IndexAny(str1, "z"))

	var t = 'd'
	fmt.Println(strings.IndexRune(str1, t))

	textString := []string{"ABC", "JKL", "XYZ"}
	fmt.Println(strings.Join(textString, "-"))

	fmt.Println(strings.LastIndex("japanese", "J"))

	MyString := "INDIA "
	repString := strings.Repeat(MyString, 5)
	fmt.Println(repString)

	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", 2))

	var S1 = strings.Split("abacadaeaf", "a")
	fmt.Println("\n", S1)

	var strr = "Australia Canada Japan"
	strr = strings.TrimPrefix(strr, "Australia")
	fmt.Println(strr)

	fmt.Println(strings.Title(str1))

	fmt.Println(strings.ToTitle(str1))

	fmt.Println(strings.ToUpper(str1))

	fmt.Println(strings.ToLower(str1))
}
