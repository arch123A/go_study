package 结构体
import "fmt"

type Drink struct {
	Name string
	Ice bool
}

func main()  {
	a:=Drink{"coffee",true,}
	b := &a
	b.Name= "tea"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(&a)
	fmt.Printf("%+v\n",a)
	fmt.Printf("%p\n",&a)
}