package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I love my work and I I I I love love love my family too too me"
	m := countmap(str)
	fmt.Println("m=", m)

}

func countmap(str string) map[string]int {
	s1 := strings.Fields(str)
	m := make(map[string]int)
	for _, value := range s1 {
		//if _ ,has :=m[value];has{
		//	m[value]++
		//}else {
		//		m[value]=1
		//}
		m[value]++

	}
	return m

}
