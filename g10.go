package main
import "fmt"

func main()  {
	china := 90
	math := 89
	english:=69
	sum :=china+math+english
	average:=float64(sum)/3
	fmt.Printf("总分是%d,平均分是%.2f", sum, average)

	fmt.Print()
}
