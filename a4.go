package main
import "fmt"

func main()  {
	a := 0
	fmt.Print("请输入a的值：")
	fmt.Scanf("%d",&a)
	fmt.Printf("a=%d\n", a)
	fmt.Print("a的地址：")
	fmt.Print(&a)
	fmt.Println("\nprintf打印")
	fmt.Printf("%p",&a)



}