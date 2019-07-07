package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)
func main() {
	s := "HELLO EVERYONE"
	s1 :=stringutil.Reverse(s)
	fmt.Println(s1)
}
