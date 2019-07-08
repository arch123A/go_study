package main

import "fmt"

func maopiao(s []int) {
	len_s := len(s)
	for i := 0; i < len_s-1; i++ {
		for j := 0; j < len_s-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}

	}
}
func main() {
	s1 := []int{9, 8, 7, 6, 5, 0, 4, 3, 2, 1}
	maopiao(s1)
	fmt.Println(s1)

}
