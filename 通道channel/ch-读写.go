package main

import (
	"fmt"
	"math/rand"
	"time"
)

func r(ch chan int) {

	n := <-ch
	fmt.Println("--随机读到的数是：", n)

}
func w(ch chan int) {
	n := rand.Intn(1000)
	ch <- n
	fmt.Println("随机写入的数字是：", n)

}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go w(ch)

	}
	for j := 0; j < 5; j++ {
		go r(ch)

	}

	for {

	}
}
