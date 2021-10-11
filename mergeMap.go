package main

import "fmt"

func main() {
	map1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(map1)

	map2 := map[string]int{
		"d": 4,
		"e": 5,
		"a": 1,
		"c": 3,
	}
	for key, value := range map2 {
		map1[key] = value
	}
	fmt.Println(map1)
}
