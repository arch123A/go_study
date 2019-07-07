package main

import "fmt"

func half(n int) (int ,error) {
	if n%2!=0{
		return -1,fmt.Errorf("cannot half %v",n)
	}
	return n/2,nil

}
func main() {
	n,r:="nnnn", "rrrrr"
	err := fmt.Errorf("the %v %v quit",n, r)
	if err!=nil{
		fmt.Println(err)
	}

	a,err:=half(3)
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println(a)
	}

	if err != nil {
		fmt.Println(err)
	}


	panic("oh,panic,panic!")
	fmt.Println("hhhhhhh")
}
