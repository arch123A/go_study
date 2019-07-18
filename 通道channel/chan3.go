package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(time.Second * 2)
	for {
		c <- "pinging"
		<-t.C
	}
}
func main() {
	m := make(chan string)
	go pinger(m)
	for i := 0; i < 5; i++ {

		msg := <-m
		fmt.Println(msg)
		fmt.Println()

	}

}
