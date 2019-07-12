package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file,err :=os.Create("/Users/yu/go/src/t1/test/b.txt")
	if err!=nil{
		fmt.Println(err)
	}
	defer file.Close()


	n1,err:=file.WriteString("aaaaaaaaaaa\n")
	n2,err:=file.WriteString("bbbbbbbbbbbb\n")
	n3,err:=file.WriteString("cccccccccc\n")
	n,err:=file.Seek(0,2)
	n4,err:=file.Seek(0,io.SeekStart)  //0,1,2分别代表文件的开始，当前位置 ，末尾
	fmt.Println(n,err,n1+n2+n3,n4)
	n5,err :=file.WriteAt([]byte("mmmmmmm\n"),n)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(n5)



}
