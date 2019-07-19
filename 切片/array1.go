package main

import "fmt"

func main() {
	//var c []string
	c := make([]string, 2)
	c[0] = "aaaaaa"
	c[1] = "bbbbb"
	c = append(c, "ccccc", "ddddd", "eeeee")
	c = append(c[:2], c[3:]...)
	//fmt.Println(c[0])
	//fmt.Println(c[1])
	d := make([]string, 4)
	copy(d, c)
	c[2] = "222222"

	fmt.Println(c)
	fmt.Println("d=", d)
}
