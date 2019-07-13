package main

import (
	"fmt"
	"strings"
)

func main() {
	var st1 string = "Hello ,go"
	a := strings.Contains(st1, "go1")
	fmt.Println(a)
	b := strings.Index(st1, "go") //返回子字符串所在的序号，位置从零开始算起
	fmt.Println(b)
	s1 := []string{"aaa", "bbb", "cccc", "ddd"}
	c := strings.Join(s1, "|") //连接字符串
	fmt.Println(c)
	d := strings.Repeat("aac", 3) //字符串重复N次
	fmt.Println(d)
	e := strings.Replace(st1, "o", "w", -1) //最后一个参数代表替换的次数，负数表示全部替换。
	fmt.Println(e)
	f := "aaaa,bbbb,cccc,dddd"
	g := strings.Split(f, ",")
	fmt.Println(g)

}
