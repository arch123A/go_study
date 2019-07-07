package main
import "fmt"

type Triangle struct {
	base float64
	heigh float64
}

func (m *Triangle) area() float64 {
	return m.base*m.heigh/2

}

func (m *Triangle) ChangeBase(a float64)  {

	m.base=a
}


func main()  {
	m:=Triangle{0.3, 0.1,}
	fmt.Println(m.area())
	m.ChangeBase(9)
	fmt.Println(m.base)
}