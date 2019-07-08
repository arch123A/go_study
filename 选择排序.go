package main

import "fmt"

func choice(s []int) {
	lenS := len(s)
	var min, minIndex int
	for i := 0; i < lenS; i++ {
		min = s[i]
		minIndex = i
		for j := i; j < lenS; j++ {
			if s[j] < min {
				minIndex = j
				min = s[j]
			}

		}
		if minIndex != i {
			s[i], s[minIndex] = s[minIndex], s[i]
		}

	}

}
func main() {

	s1 := []int{9, 8, 7, 6, 8, 95, 4, 3, 2, 1, 10, 11}
	choice(s1)

	fmt.Println(s1)
}
