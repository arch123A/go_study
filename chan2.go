package main

import (
	"fmt"
	"time"
)

func received(c chan string) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	c := make(chan string, 2)
	c <- "message111"
	c <- "message2222"
	close(c)
	fmt.Println("close channel c")
	time.Sleep(time.Second*1)
	received(c)


}
