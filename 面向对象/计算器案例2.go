package main

import "fmt"

type Op interface {
	operation() int
	setdata(i ...interface{}) bool
}

type Object struct{
    n1 int
    n2 int

}


type add struct{
    Object
}
func (s *add) operation()  int{
	return s.n1+s.n2
}

//初始化赋值的方法
func (s *add) setdata(i ...interface{})  bool{

	var r bool=true
	if len(i)!=2{
		fmt.Println("参数个数错误！")
		r=false

	}
	for index, value := range i {
		v,ok:=value.(int)
		if !ok{
			fmt.Println("参数类型错误！")
			r=false

		}else {
			if index==0{
				s.n1=v
			}else {
				s.n2=v
			}

		}
	}
	return r


}




type sub struct {
	Object
}
func (s *sub) operation()  int{
	return s.n1-s.n2
}
func (s *sub) setdata(i ...interface{})  bool{
	//初始化赋值的方法
	var r bool=true
	if len(i)!=2{
		fmt.Println("参数个数错误！")
		r=false

	}
	for index, value := range i {
		v,ok:=value.(int)
		if !ok{
			fmt.Println("参数类型错误！")
			r=false

		}else {
			if index==0{
				s.n1=v
			}else {
				s.n2=v
			}

		}
	}
	return r


}

func boot2(s Op)  int{
	return s.operation()

}

type third struct{
    
}

func (s *third)result(c string)  Op{
	switch c {
	case "+":
		a:=new(add)
		return a
	case "-":
		b:=new(sub)
		return b
	default:
		fmt.Println("输入错误内容！")
		return nil
		
	}
}


func main() {

	//var a add
	//var b sub
	var c third
	r:=c.result("-")
	b1:=r.setdata(10, 5)


	if b1{
		fmt.Println(boot2(r))
	}


}
