package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) show() {
	fmt.Printf("我叫%s，今年%d了。\n", p.name, p.age)

}

func main() {

	p := Person{"张三", 25}
	f := p.show //方法值
	f()

	f2 := (*Person).show //方法表达式
	f2(&p)

}
