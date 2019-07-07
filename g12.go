package main
import "fmt"

func main()  {
    var year int
    fmt.Println("请输入年份：")
    fmt.Scan(&year)
    //逻辑与的优先级大于逻辑或，帮先算后面两项
    if	year%400==0 || year%4==0 && year%100 != 0 {
		fmt.Printf("%d年是闰年", year)
	} else {
		fmt.Printf("%d年不是闰年", year)

	}

}