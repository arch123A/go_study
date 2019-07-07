package main

import "fmt"

func main() {
	var score int
	fmt.Println("请输入你的成绩")
	fmt.Scan(&score)
	if score >= 90 && score <= 100 {
		fmt.Println("你的成绩是A")
	} else if score >= 80 {
		fmt.Println("你的成绩是B")
	} else if score >= 70 {
		fmt.Println("你的成绩是C")
	} else if score >= 60 {
		fmt.Println("你的成绩是D")
	} else {
		fmt.Println("你的成绩不及格")
	}

}
