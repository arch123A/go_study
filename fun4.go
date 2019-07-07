package main

import "fmt"

func h10(n int) int {

	if n <= 0 {
		return 1
	}
	fmt.Println("n=", n)
	return h10(n - 1)

}

func main() {
	//h10(10)
	//fmt.Println()
	f2c := func(f float64) float64 {
		return 5 * (f - 32) / 9
	}

	fmt.Println(f2c(212))
	i:=0
	for i<10{
		fmt.Println(i)
		i++
	}
}
