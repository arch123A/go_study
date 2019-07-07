package main
import "fmt"

func main()  {
    var num int
	fmt.Println("请输入一个整数：")
    fmt.Scan(&num)
    if num%2==0{
    	fmt.Printf("%d是一个偶数。", num)

	} else {
		fmt.Printf("%d是一个奇数。", num)
	}

}