package main

import (
	"fmt"
	"net"
)

/*
TCP-CS客户端：

1.  conn, err := net.Dial("TCP", 服务器的IP+port)

2.  写数据给 服务器 conn.Write()

3.  读取服务器回发的 数据 conn.Read()

4.  conn.Close()

*/

func main() {
	//与主机建立起链接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Println("客户端已经与主机端建立 连接！")
	//创建切片buf用来读取主机内容，亦是s用来保存输入的内容
	buf := make([]byte, 1024)
	var s string
	for {

		fmt.Println("请输入需要发送的内容：")
		fmt.Scan(&s)
		//向主机发送内容
		conn.Write([]byte(s))

		//接受主机发送的内容并打印
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("主机端发来消息", string(buf[:n]))
	}

}
