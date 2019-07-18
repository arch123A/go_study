package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

func productor(ch chan<- int, i int) {
	for {
		cond.L.Lock()
		for len(ch) == 3 {
			cond.Wait()
		}
		n := rand.Intn(1000)
		ch <- n
		fmt.Printf("--%d、生产的数是%d\n", i, n)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 300)

	}

}
func consumer(ch <-chan int, i int) {
	for {
		cond.L.Lock()
		for len(ch) == 0 {
			cond.Wait()
		}

		n := <-ch
		fmt.Printf("%d、消费的数是%d\n", i, n)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 300)

	}

}

func main() {
	ch := make(chan int, 3)
	quit := make(chan bool)
	//rand.Seed(time.Now().UnixNano())

	// 指定条件变量 使用的锁
	cond.L = new(sync.Mutex) // 互斥锁 初值 0 ， 未加锁状态

	for i := 0; i < 5; i++ {
		go productor(ch, i)
		go consumer(ch, i)

	}
	/*for j := 0; j < 5; j++ {
		go consumer(ch, j)

	}*/

	<-quit
	//	for {
	//		;
	//	}
}
