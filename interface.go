package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
	Talk() string
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error  {
	return nil

}
func (a *T850) Talk() string  {
	return "我正在使用T850"

}

type R2d2 struct {
	Broken bool
}

func (a *R2d2) PowerOn() error  {
	if a.Broken{
		return errors.New("R2d2 is broken")

	}else {
		return nil
	}
}
func (a *R2d2) Talk() string  {
	return "我正在使用R2d2"

}


func Boot(r Robot) error {
	return r.PowerOn()
}
func Rtalk(r Robot) string {
	return r.Talk()
}


func main()  {
	t := T850{
		"The Terminator",
	}
	r := R2d2{true,}
	err := Boot(&r)
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("Robot is PowerOn")

	}
	fmt.Println(t.Talk())

	fmt.Println("T850")
	err=Boot(&t)
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("Robot is PowerOn")

	}
	fmt.Println(r.Talk())

}