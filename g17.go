package main
import "fmt"

func main()  {
	var acount,sit_count int=3,1
	//acount,sit_count
	if acount>2{
		fmt.Println("余额充足，可以上车。")
		if sit_count>0{
			fmt.Println("有座位，请坐下")
		}else{
			fmt.Println("没有座位，请站着。")
		}
	}else{
		fmt.Println("余额不足，不可以上车。")
	}

}