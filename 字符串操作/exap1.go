package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2, str3 string
	//fmt.Println("please input:")
	//fmt.Scan(&str1)
	//str2= "年 月 日"
	//s2:=strings.Split(str2," ")
	//s1:=strings.Split(str1,"-")
	//if len(s1)==len(s2){
	//	for i:=0;i<len(s2);i++  {
	//		str3+=s1[i]+s2[i]
	//	}
	//}else{
	//	fmt.Println("输入错误！")
	//}
	//
	//
	//fmt.Println(str3)

	str1 = "老王很邪恶"
	if strings.Contains(str1, "邪恶") {

		str2 = strings.Replace(str1, "邪恶", "**", -1)
	}
	fmt.Println(str2, str3)
}
