package main

import (
	"fmt"
	"os"
)

func main() {
	file,err :=os.OpenFile("/Users/yu/go/src/t1/test/b.txt",os.O_RDWR, 6)
	if err!=nil{
		fmt.Println(err)
	}
	defer file.Close()
	n1,err:=file.Seek(0,2)
	n,err :=file.WriteAt([]byte("nnnnnnnn\n"),n1)
	if err!=nil{
		fmt.Println(err)

	}
	fmt.Println(n)


}
