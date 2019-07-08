package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{"aaa": 1, "bbb": 2, "ccc": 3}
	value, ok := m["ccc"]
	if ok {
		fmt.Println(value)
		fmt.Println("存在！")

	} else {
		fmt.Println("m[\"ccc\"]", "不存在 ")
	}
}
