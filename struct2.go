package main
import "fmt"

type Hero struct {
	name string
	age int
	address Address
}

type Address struct {
	num int
	street string
	city string
}



func main()  {
	//结构体的嵌套
    m1:= Hero{"greenspider", 23, Address{11, "xxxxli", "beijing",},}


		fmt.Println(m1)
	}