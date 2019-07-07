package main
import "fmt"

func main()  {
    const F1 float64=32.33
    const PI=3.14
	//f1=22.33,常量不允许修改值,名子建议大写
	fmt.Print(F1)
    fmt.Println()
	fmt.Print(PI)
	//fmt.Printf("%p", &F1) // 不允许打印常量的地址


}
