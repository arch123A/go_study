package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file1,err1:=os.Open("/Users/yu/go/src/t1/test/c.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer file1.Close()

	file2,err2:=os.Create("/Users/yu/go/src/t1/test/c2.txt	")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer file2.Close()
	buffer:=make([]byte, 1024)


	for{
		n,err3 :=file1.Read(buffer)
		if err3 == io.EOF {
			break
		}
		//n2,_:=file2.Seek(0,2)
		//file2.WriteAt(buffer[:n],n2)
		file2.Write(buffer[:n])


	}



}
