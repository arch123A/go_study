package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main()  {
    var a int=2
    var s string= "123"
    var b string= "trwue"
    c,err:=strconv.ParseBool(b)
    if err!=nil{
    	fmt.Println("出现了错误：",err)
    	fmt.Println("aaa"+"bbbb")
	}
    //转换字符串成数字
    d,err2:=strconv.ParseInt(s,0,8)
	if err2!=nil{
		fmt.Println("出现了错误：",err2)

	}else{
		fmt.Println("转换后的数值是",d)
	}



	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(err))
    fmt.Println(c,err)
}