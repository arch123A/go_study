package main

import (
	"bytes"
	"fmt"
)

func main() {
	s1 := "aaa"
	var buffer bytes.Buffer

	//for k := 1; k < 50000; k++ {
	//	s1 += strconv.Itoa(k)
	//	if k%50 == 0 {
	//		s1 += "\n"
	//	}
	//
	//}

	fmt.Println(s1)
	for i:=0;i<500;i++  {
		buffer.WriteString("windows\t")
	}
	fmt.Println(buffer.String())
	for j:=0;j<=500;j++  {
	        fmt.Println()
			
	}
}
