package main

import "fmt"

type Personer interface {
	show()
}

type Student struct {
}

func (s *Student) show() {
	fmt.Println("大家好，我是学生。")

}

type Teacher struct {
}

func (s *Teacher) show() {
	fmt.Println("大家好，我是老师。")

}

//多态的实现
func boot(p Personer) {
	p.show()

}

func main() {
	var s Student
	var t Teacher
	//var p Personer
	////p=&s
	//p.show()
	//p=&t
	//p.show()
	boot(&s)
	boot(&t)

}
