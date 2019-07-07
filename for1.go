package main

import "fmt"

func main()  {
    //s1:=[]int{1, 2, 3, 44, 55, 66,}
	//for index, s := range s1 {
	//	 defer fmt.Println("index of loop is ", index)
	//	fmt.Println("The value from the slce is ",s)
	//}
	//fmt.Println()


	var i int
	fmt.Println("请输入一个整数：")
	fmt.Scan(&i)
	//switch i {
	//case 1:
	//	fmt.Println("one")
	//case 2:
	//	fmt.Println("two")
	//case 3:
	//	fmt.Println("three")
	//
	//}

	if i>5 && i<10{
		fmt.Println("It's ok!")
	}
}
