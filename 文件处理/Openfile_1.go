package main

import (
	"fmt"

	"os"
)

func main() {
	file, err := os.OpenFile("/home/yu/go/src/test1/a.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	n1, err := file.Seek(0, 2)
	n, err := file.WriteAt([]byte("33333333333333333\n"), n1)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(n1)
	fmt.Println(n)

}
