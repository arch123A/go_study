package main
import "fmt"

func main()  {
	str := "中国人民"
	//在golang中一个汉字占三个字符
	fmt.Printf("str类型是%T", str)
	fmt.Print("\nstr的字符个数为：",len(str))
	fmt.Println("")
	i1 := 234545.444555
	fmt.Println(int(i1))



}
