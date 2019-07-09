package main

import "fmt"

func modify(sli []int) []int {
	for k := 0; k < 5; k++ {
		//sli[k]=k+1
		sli = append(sli, k+1)
	}
	return sli
}
func remove(s []int, index int) []int {
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}

func main() {
	//s := make([]int, 20)
	//s=modify(s)
	//fmt.Println(s)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s=remove(s,6)
	//fmt.Println(s)
	//s=append(s,s[2:4]...)
	//fmt.Println(s)
	for key, value := range s {
		fmt.Println(key, value)
	}
}
