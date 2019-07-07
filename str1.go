package main

import "fmt"

func main() {
	s := "aaa\n\\ \t" +
		"\\\"bb" +
		"b"
	b := `cccc\n
			ddd
			ddd
			fff
			ggg
			`

	fmt.Println(s)
	fmt.Println(b)
}
