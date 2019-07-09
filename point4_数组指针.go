package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	var p *[5]int
	p = &arr
	//[]的优先级高于*,所以 要在*p上加括号，也可以简写成p[0]
	fmt.Println((*p)[4])
	fmt.Println(p[4])
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])

	}

}
