package main

import "fmt"

func qc(s1 []string) []string {

	//返回去掉重复值的切片
	s2 := s1[:1]
	for _, v1 := range s1[1:] {
		is_in := false
		for _, v2 := range s2 {
			if v1 == v2 {
				is_in = true
				break
			}

		}
		if !is_in {
			s2 = append(s2, v1)
		}
	}
	return s2

}

func qk(s []string) []string {

	//返回去除空值的字符串切片,并且在原有切片上操作
	var i int = 0
	for _, value := range s {
		if value != "" {
			s[i] = value
			i++

		}
	}
	//s=s[:i]
	fmt.Println("qk_func:", s[:i])
	return s[:i]

}

func qk2(s []string) []string {
	//返回去除空值的字符串切片

	//s1:=[]string{}
	s1 := make([]string, 0, 1)
	for _, value := range s {
		if value != "" {
			s1 = append(s1, value)
		}
	}

	return s1

}

func main() {
	//s1 := []string{"green", "red", "yello", "red", "apple", "pink", "green", "brown", "brown"}
	s2 := []string{"green", "red", "", "yello", "red", "", "apple", "pink", "green", "brown", "brown"}
	//fmt.Println(qc(s1))
	fmt.Println(qk(s2))
	//s2=s2[:]
	fmt.Println(s2)
	fmt.Println(qk2(s2))
}
