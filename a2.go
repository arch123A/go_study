package main

import "fmt"

func main()  {
	var a1,a2 int
	a1=99
	var a3 int=100
	var a4 int=200
	a4,a3=a3,a4
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(">>>>>>>>>>>>")
	fmt.Println(a3,a4)
	//fmt.Println(a4)


}

