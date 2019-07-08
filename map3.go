package main

import (
	"fmt"
)

func main() {
	s1 := "hello,everyone"
	var m map[byte]int
	//m:=make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		ch := s1[i]
		m[ch] = m[ch] + 1
	}
	fmt.Println(m)
	for key, value := range m {

		fmt.Printf("%c:%d\n", key, value)
	}
}
