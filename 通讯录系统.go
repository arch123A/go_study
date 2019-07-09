package main

import "fmt"

type Person struct {
	name  string
	phone map[string]int
}

var tx1 map[string]int = map[string]int{"办公室": 8888, "家里": 9999}

var personList []Person = []Person{
	{"张三", tx1},
	{"李四", map[string]int{"办公室": 6666, "家里": 7777}},
}

func main() {
	switch_1()

	//showPerson()
}

//显示操作提示，做不同的操作
func switch_1() {
	var require int //需要进行的操作名称（1，2，3，4）
	fmt.Println("增加联系人请按1")
	fmt.Println("查找联系人请按2")
	fmt.Println("修改联系人请按3")
	fmt.Println("删除联系人请按4")
	fmt.Print("请输入：")
	fmt.Scan(&require)
	switch require {
	case 1:
		addPerson()

	case 2:
		index := findPerson()
		if index != -1 {
			fmt.Println("姓名：", personList[index].name)
			for key, value := range personList[index].phone {
				fmt.Printf("%s的电话是：%d\n", key, value)
			}
		} else {
			fmt.Println("查无此人。")

		}

	case 3:
		fmt.Println("修改联系人")
	case 4:

		index := findPerson()
		if index != -1 {
			personList = append(personList[:index], personList[index+1:]...)
			fmt.Println("删除成功！")
			showPerson()
		} else {
			fmt.Println("查无此人。")
		}

	}

}

func addPerson() {
	var name, type_phone string
	var phone int
	var p1 Person
	var introduce string

	fmt.Println("请输入姓名：")
	fmt.Scan(&name)
	for {
		fmt.Println("请输入电话号码类型：")
		fmt.Scan(&type_phone)
		fmt.Println("请输入电话号码：")
		fmt.Scan(&phone)
		p1.name = name
		p1.phone = map[string]int{type_phone: phone}
		//personList=append(personList,Person{name, map[string]int{type_phone:phone}})
		personList = append(personList, p1)
		fmt.Println("按q退出，按其他键，继续输入电话号码。")
		fmt.Scan(&introduce)
		if introduce == "q" || introduce == "Q" {
			break

		}

	}

}

func findPerson() int {
	var name string
	var index int = -1
	fmt.Println("请输入姓名：")
	fmt.Scan(&name)

	for key, value := range personList {
		if value.name == name {
			return key

		}

	}
	return index

}

func showPerson() {
	for _, person := range personList {
		fmt.Printf("姓名是:%s\n", person.name)
		for key, value := range person.phone {
			fmt.Printf("%s的电话号码是:%d\n", key, value)

		}

	}
}
