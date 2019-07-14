package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/home/yu/go/src/test1/a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buffer1 := make([]byte, 10)
	for {

		n, err := file.Read(buffer1)

		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		fmt.Print(string(buffer1[:n]))

	}

}
