package main
import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("请输入你的姓名：")
	fmt.Scan(&name)

	fmt.Println("请输入你的年龄：")
	fmt.Scan(&age)
	fmt.Printf("姓名是%s,年龄是%d",name,age)
}
