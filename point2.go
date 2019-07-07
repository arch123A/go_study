package main

import "fmt"

func t1(n ,m int) {
	fmt.Println(&n)
	fmt.Println(&m)
}

//判断是否是偶数
func isEvent(a int) bool{
	return a%2==0

}
func test1() (int ,string){
	a := 1
	b := "sssss"
	return a, b

}

func sum(n ...int) int {
	s:=0
	for k := 0; k < len(n); k++ {
		s+=n[k]

	}
	return s
}

func s2(n... int) int {
	s := 0
	for _, nn := range n {
		s+=nn
	}
	return s
}

func main() {
//defer采用后进先出的策略
//	a:=2
//
//	fmt.Println(&a)
//	fmt.Println(isEvent(12),isEvent(13))
//	fmt.Printf("\n%v",isEvent(9))

	//c,d:=test1()
	//fmt.Printf("c:%v,d:%v",c, d)

	sum_ :=sum(1,2,3,45)
	ss2:=s2(1,2,3,4,4,4,4,4,45,5,55,5,0)
	fmt.Printf("sum=%v\n",sum_)
	fmt.Printf("ss2=%v", ss2)




}
