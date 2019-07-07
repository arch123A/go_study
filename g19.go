package main
import "fmt"

func main()  {
    var account,passwd string
	fmt.Println("请输入用户名")
    fmt.Scan(&account)
	fmt.Println("请输入密码")
	fmt.Scan(&passwd)
    if account== "admin" {
    	if passwd== "8888" {
    		fmt.Println("帐号和密码正确。")
		}else {
			fmt.Println("帐号正确，密码错误。")
		}
	}else{
		if passwd== "8888" {
			fmt.Println("帐号错误，密码正确。")
		}else {
			fmt.Println("帐号错误，密码错误。")
		}

	}

}