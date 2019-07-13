package main

import "fmt"

type Op interface {
	operation() int
}

type Object struct {
	n1 int
	n2 int
}

type add struct {
	Object
}

func (s *add) operation() int {
	return s.n1 + s.n2
}

type sub struct {
	Object
}

func (s *sub) operation() int {
	return s.n1 - s.n2
}

func boot2(s Op) int {
	return s.operation()

}

type third struct {
}

func (s *third) result(a, b int, c string) int {
	switch c {
	case "+":
		add1 := add{Object{a, b}}
		return add1.operation()
	case "-":
		sub1 := sub{Object{a, b}}
		return sub1.operation()
	default:
		fmt.Println("输入错误内容！")
		return 0

	}
}

func main() {
	/*
		//var a add=add{Object{1,2}}
		//var b sub=sub{Object{3,1}}
		////var s Op
		//c:=boot2(&a)
		//fmt.Println(c)
		//c=boot2(&b)
		//fmt.Println(c)

		//s=&a
		//fmt.Println(s.operation())
		//
		//s=&b
		//fmt.Println(s.operation())
	*/

	var p third
	r := p.result(10, 5, "-")
	fmt.Println(r)

}
