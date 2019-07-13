package main

import "fmt"

//切片去重
func noempty1(s []string) []string {
	result := s[:0]
	for _, value := range s {
		if value != "" {
			result = append(result, value)
		}

	}
	return result
}

func noempty2(s []string) []string {
	i := 0
	for _, value := range s {
		if value != "" {
			s[i] = value
			i++
		}
	}
	return s[:i]

}

func main() {
	data := []string{"red", "", "black", "", "", "pink", "blue"}
	d2 := noempty2(data)
	fmt.Printf("%v", d2)
}
