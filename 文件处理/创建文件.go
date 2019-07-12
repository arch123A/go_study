package main

import (
	"fmt"
	"os"
)

func main()  {
    file,err :=os.Create("/Users/yu/go/src/t1/test/a.txt")
    if err!=nil{
    	fmt.Println(err)

	}
    defer file.Close()


    n1,err:=file.WriteString("cccccccccccccc\n")
	n2,err:=file.Write([]byte("dddddddddddd\n"))
	n3,err:=file.WriteString("eeeeeeeeeeeeeeee\n")
    fmt.Println(n1,n2,n3)


}