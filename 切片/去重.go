package main

import (
	"fmt"
)

//切片的去重
func qc(s []string) []string {
	if len(s) == 0 {
		return nil
	}
	s1 := s[:1]
	for _, value := range s {
		is_re := false
		for i := 0; i < len(s1); i++ {
			if value == s[i] {
				is_re = true
				break
			}

		}
		if !is_re {
			s1 = append(s1, value)
		}
	}
	return s1

}

func main() {
	//data := []string{"red", "black", "red", "yellow", "yellow", "pink", "blue", "pink", "blue"}
	s2 := []string{}
	s := qc(s2)
	fmt.Printf("%v\n", s)
}
