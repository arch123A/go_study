package main

import "fmt"

func main() {
	t_shirt := 35
	trousers := 120
	sum := t_shirt*3 + 2*trousers
	account:=float64(sum)*0.88
	result:=int(account)

	fmt.Printf("折扣后的价格是%.2f\n", account)
	fmt.Printf("需要付的钱是%d\n", result)

	fmt.Print()
}
