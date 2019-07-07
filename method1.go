package main

import (
	"fmt"
	"strconv"
	
)

//Movie 
type Movie struct {
	Name string
	rating float64
}

//创建方法
func (m *Movie)summary() string {
	r:=strconv.FormatFloat(m.rating,'f',1,32)
	return m.Name+","+r

}


func main()  {
	a := Movie{"光明皇帝",3	}
	b := a.summary()
	fmt.Println(b)

    
	fmt.Println()
}