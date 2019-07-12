package main

import (
	"fmt"
	"os"
)

func main() {
	file,err :=os.Open("/Users/yu/go/src/t1/test/b.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buffer:=make([]byte, 1024)

	n,err :=file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buffer[:n]))


}
