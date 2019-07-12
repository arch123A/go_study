package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file,err :=os.Open("/Users/yu/go/src/t1/test/c.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buffer1 :=make([]byte,10)
	for{

		n,err :=file.Read(buffer1)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				break
			}
		}
		fmt.Print(string(buffer1[:n]))


	}


}