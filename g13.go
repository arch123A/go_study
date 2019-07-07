package main
import "fmt"

func main()  {
    var chinese,math int
    fmt.Println("请输入语文成绩：")
    fmt.Scan(&chinese)
	fmt.Println("请输入数学成绩：")
	fmt.Scan(&math)
    if chinese>70 && math==100{
    	fmt.Println("成绩不错，奖励一百元。")
	}else {
		fmt.Println("成绩一般，还要努力。")

	}

	fmt.Print()
}