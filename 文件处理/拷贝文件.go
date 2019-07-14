package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file1, err1 := os.Open("/home/yu/go/src/test1/c.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer file1.Close()

	file2, err2 := os.Create("/home/yu/go/src/test1/c2.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer file2.Close()
	buffer := make([]byte, 1024)

	for {
		n, err3 := file1.Read(buffer)
		if err3 == io.EOF {
			fmt.Println("finished!")
			break
		} else if err3 != nil {
			fmt.Println(err3)
			break
		}
		//n2,_:=file2.Seek(0,2)
		//file2.WriteAt(buffer[:n],n2)
		file2.Write(buffer[:n])

	}

}
