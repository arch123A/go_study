package main

import "fmt"

type Student struct {
	name string
	age  int
	id   int
}
type Movie struct {
	name   string
	rating float32
}

func main() {
	//m1:=Student{
	//	name: "张三",
	//	age:29,
	//	id:2,
	//
	//}
	//fmt.Println(m1)
	//fmt.Println(m1.name,m1.age)

	var m Movie
	fmt.Println(m)
	fmt.Printf("m:%+v\n", m)
	m.name = "美丽的生活"
	m.rating = 2
	fmt.Printf("m:%+v\n", m)

	b := Movie{"小小人生", 23}
	fmt.Println(b)
}
