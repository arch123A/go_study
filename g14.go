package main
import "fmt"

func main()  {
    var age int
    fmt.Println("请输入你的年龄：")
    fmt.Scan(&age)
    if age>18 {
    	fmt.Println("已经满18岁，可以进网吧。")
	}else
    {
    	fmt.Println("没满18岁，不能进网吧。")
	}
	fmt.Print()
}