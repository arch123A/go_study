package main

import "fmt"

func main() {
//defer采用后进先出的策略
	defer fmt.Println("11111")
	 fmt.Println("222222")
	 fmt.Println("33333")


}
