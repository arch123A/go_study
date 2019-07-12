package main

import "fmt"

type operation struct {
}

func (s *operation) ys(a, b int, c string) int {
	var result int
	switch c {
	case "+":
		result = a + b
	case "-":
		result = a - b

	default:
		fmt.Println("错误！！")
		result = 0
	}
	return result
}

func main() {
	var a operation
	b := a.ys(5, 2, "+")
	fmt.Println(b)
}
