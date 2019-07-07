package main

import "fmt"

func jc(n int) int{
	if n==1{
		return 1
	}else{
		return n*jc(n-1)
	}

}

func jc2(n int) int {
	var result int=1
	for i := 1; i <= n; i++ {
		result*=i
	}
	return result
}

func main()  {
	//
	//fmt.Println(jc(10))
	//fmt.Println(jc2(10))
	var a[5] int =[5]int{1, 2, 3, 4}
	fmt.Println(a[4])
}