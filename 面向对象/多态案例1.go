package main

import "fmt"

//接口类型名以er结尾
type Storager interface {
	Read()
	Write()
}

type Udisk struct {
}

func (s *Udisk) Read() {
	fmt.Println("正在读取u盘内容。")
}

func (s *Udisk) Write() {
	fmt.Println("正在写入u盘。")

}

type Disk struct {
}

func (s *Disk) Read() {
	fmt.Println("正在读取硬盘内容。")
}

func (s *Disk) Write() {
	fmt.Println("正在写入硬盘。")

}

func boot1(a Storager) {
	a.Read()
	a.Write()

}

func main() {
	var a Udisk
	var b Disk
	boot1(&a)
	boot1(&b)

}
