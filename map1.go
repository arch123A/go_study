package main
import "fmt"

func main()  {
    var m1 map[int]string= map[int]string{1:"aaa",2:"bbbb"}

    m2:=map[int]string{}
	//m2[2]="aaaa"
	fmt.Println(m1== nil)
    fmt.Println(m2== nil)
	m1[3]="cccc"
	fmt.Println("m1=",m1)

	fmt.Println(m2)


	m3:=make(map [int]string)
	m3[3]="bbbbb"
	fmt.Println(m3)
}
