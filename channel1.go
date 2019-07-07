package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string)  {
	time.Sleep(time.Second*2)
	fmt.Println("slowfunc finished")
	c <- "from slowfunc"
}

func main() {
	c := make(chan string)
	go slowFunc(c)

	msg:=<-c
	fmt.Println(msg)
	fmt.Println("mani filished")
	//time.Sleep(time.Second*3)

}
