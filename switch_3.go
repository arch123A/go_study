package main

import "fmt"

func main() {
	var score string
	var salary int=5000
	fmt.Println("请输入你的评级：")
	fmt.Scan(&score)

	switch score {
	case "a":
		fmt.Printf("你的工资是%d\n",salary+500)
	case "b":
		fmt.Printf("你的工资是%d\n",salary+200)
	case "c":
		fmt.Printf("你的工资是%d\n",salary)
	case "d":
		fmt.Printf("你的工资是%d\n",salary-200)
	case "e":
		fmt.Printf("你的工资是%d\n",salary-500)
	default:
		fmt.Println("输入错误")


	}
}

