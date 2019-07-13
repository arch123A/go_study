package main

import "fmt"

func main() {
	var i interface{}
	i= "aaaa"
	fmt.Println(i)
	i=2345

	fmt.Println(i)

	var b []interface{} //空接口切片
	b=append(b,11,22,33,44,"ccc","dddd", "asss")
	fmt.Println(b)
	for index, value := range b {
		fmt.Println(index+1,value)
	}


}
