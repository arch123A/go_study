package main

import "fmt"

func main() {
	var p *string
	p = new(string)
	*p = "hello,wwwwww98uuuuuu99999999999uuû"
	fmt.Println(*p)
}
