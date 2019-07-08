package main

import "fmt"

func quickSort(s []int, left int, right int) {

	pivot := s[(left+right)/2]
	i, j := left, right
	for i <= j {

		for s[i] < pivot {

			i++

		}
		for s[j] > pivot {

			j--
		}
		if i <= j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	if left < j {
		quickSort(s, left, j)
	}
	if i < right {
		quickSort(s, i, right)
	}

}
func main() {
	s1 := []int{9, 8, 7, 6, 8, 95, 1, 4, 3, 2, 1, 10, 11}
	s2 := []int{3, 2, 1}
	quickSort(s2, 0, len(s2)-1)
	quickSort(s1, 0, len(s1)-1)

	fmt.Println(s2)
	fmt.Println(s1)
}
