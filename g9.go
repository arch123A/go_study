package main
import "fmt"

func main()  {
    const PI =3.14
    var r float64
	fmt.Println("请输入半径：")
    fmt.Scan(&r)

	area := PI*r*r
	fmt.Printf("面积是%.2f\n", area)
	fmt.Printf("周长是%.2f", 2*PI*r)

	fmt.Print()
}
