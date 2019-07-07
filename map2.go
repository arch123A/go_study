package main
import "fmt"

func main()  {
    d1:=make(map[string]int)
    d1["aa"]=1
    d1["bb"]=2
    d1["cc"]=3
	delete(d1,"aa")
	fmt.Println(d1)

    //d2:=make(map[string]string)
    var d3 map[string]string = map[string]string{"p": "段落", "img": "图像", "h1": "一级标题", "h2": "二级标题"}
    fmt.Println(d3)
	for key, value := range d3 {
		fmt.Println("key:",key)
		fmt.Println("value:",value)
	}
}