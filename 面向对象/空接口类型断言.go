package main

import "fmt"

func main() {
	var i interface{}
	i= "sdssdds"
	v,ok:=i.(string)
	if ok  {
		fmt.Println(v)

	}else{
		fmt.Println("类型断言错误。")
	}

}
