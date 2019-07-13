package main

import "fmt"

func remove(s []int, n int) []int {
	copy(s[n:], s[n+1:])
	return s[:len(s)-1]
}

func main() {

	s := []int{1, 2, 3, 4, 5, 65, 7, 8, 9, 10}
	s2 := remove(s, 5)
	fmt.Println(s2)

}
