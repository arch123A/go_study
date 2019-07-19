package main

import "fmt"

var big []int = []int{1, 3, 5, 7, 8, 10, 12}
var small []int = []int{4, 5, 9, 11}

func is_in(s1 []int, mon int) bool {

	for _, value := range s1 {
		if mon == value {
			return true
		}

	}
	return false

}

//判断年的月份有多少天
func main() {
	for {
		var year, month int = 2008, 12
		var day int
		fmt.Println("请输入年份：")
		fmt.Scan(&year)
		fmt.Println("请输入月份：")
		fmt.Scan(&month)
		var is_run bool
		if year%400 == 0 || year%4 == 0 && year%100 != 0 {
			is_run = true
		} else {
			is_run = false
		}

		switch {
		case is_in(big, month):
			day = 31
		case is_in(small, month):
			day = 30
		case month == 2:
			if is_run {
				//fmt.Printf("%d年%d月有29天\n", year, month)
				day = 29

			} else {
				day = 28

			}
		default:
			fmt.Println("你的输入有误！")
			continue
		}

		fmt.Printf("%d年%d月有%d天\n", year, month, day)
	}
}
