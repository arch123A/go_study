package main

import (
	"bufio"
	"fmt"
	"io"

	"os"
)

func main() {
	file, err := os.OpenFile("/home/yu/go/src/test1/a.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println("Successful")
	//创建带有缓冲区的reader
	read := bufio.NewReader(file)

	for {
		b1, err := read.ReadBytes('\n')

		if err != nil && err == io.EOF {
			fmt.Println("finished!")
			break
		} else if err != nil {
			fmt.Println(err)
			break

		}
		fmt.Println(string(b1))
	}

}
