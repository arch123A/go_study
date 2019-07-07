package main
import "fmt"

func main()  {
    /*var age int
    fmt.Println("请输入年龄：")
    fmt.Scan(&age)
	fmt.Print("age=", age)*/

    //var f1 float64
    f1 := 2.88
    fmt.Printf("\n%.2f",f1)
    fmt.Printf("\n数据类型是%T",f1)
    i := 24
    fmt.Printf("\n%d",i)
    fmt.Printf("\ni数据类型是%T",i)

    var bi bool
    bi=true
    fmt.Printf("\nbi=%t", bi)
    fmt.Printf("\nbi数据类型是：%T", bi)

    var ch byte
    //var i int
    //ch= 97
    for i=0; i < 26; i++ {
        ch= byte(i + 97)
        fmt.Printf("\nch=%c", ch)
    }
    fmt.Printf("\nch=%c", ch)
}
