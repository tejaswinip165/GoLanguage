package main

import "fmt"

func main() {
	map1 := map[int]string{

		100: "mr.x",
		101: "mr.y",
		102: "mr.z",
		103: "mr.x",
		104: "mr.z",
	}
	fmt.Println(map1)

	for k := range map1 {
		if k == 104 {
			map1[104] = "mr.l"
			fmt.Println(map1)

		}
	}
}
