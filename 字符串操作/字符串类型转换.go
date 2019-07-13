package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 字符串转换成 bool
	a, err := strconv.ParseBool("F")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

	//将bool类型转换成字符串
	b1 := strconv.FormatBool(false)
	fmt.Println(b1)

	//将整型转换成字符串，和将字符串转换成整型
	i1, err := strconv.Atoi("2345")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i1)

	str1 := strconv.Itoa(45679)
	fmt.Println(str1)

}
