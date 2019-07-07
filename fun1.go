package main

import "fmt"

func register() {
	var account, pwd, email string
	account = ""
	pwd = "123"
	email = "aa@123.com"
	b := check(account, pwd, email)
	if b{
		fmt.Println("你的注册信息无误，已经通过验证")
	}else{

		fmt.Println("你的注册信息不完整")
	}


}

func check(a string, p string, e string) bool{
	if a == "" || p == "" || e == ""{
		return false

	}else{
		return true

	}

}

func main() {
	register()
	//fmt.Println()
}
