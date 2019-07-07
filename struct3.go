package main

import (
	"fmt"
	"reflect"
)
type Alarm struct {
	Time string
	Sound string
}

//使用构造函数自定义默认值
func NewAlarm(ini string) Alarm{
	a:=Alarm{
		ini, "keoo",

	}
	return a
}
func main()  {
	b := NewAlarm("8:00")
	fmt.Println(NewAlarm("7:00"))
	fmt.Printf("%T\n",b)
	fmt.Println(reflect.TypeOf(b))
}