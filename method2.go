package main
import ("fmt"
	"math")



type sphere struct {
	r float64
}

func (m *sphere) area() float64 {
	return float64(4)*math.Pi*m.r*m.r
}
func (m *sphere) volumn() float64 {
	tem:=math.Pow(m.r,3)
	return tem*4/3*math.Pi
}


func main()  {
	a := sphere{5}
	fmt.Println(a.area())
	fmt.Println(a.volumn())
}