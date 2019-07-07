package main
import "fmt"

func main()  {
    var date int
	fmt.Println("请输入数字（1-7）")
    fmt.Scan(&date)
	switch date {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("你输入的有误。")

	}
}