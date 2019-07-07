package main
import "fmt"


func anotherfunc(b func() string) string{
	return b()
}
func main()  {
    a:= func() string{
		return "aaaaa"
	}
    anotherfunc(a)
	fmt.Println(a())
	fmt.Println(anotherfunc(a))
}
