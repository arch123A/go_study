package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器端等待客户端连接")
	for {
		//监听客户端
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handledata(conn)
	}

}

//读取客户端信息,把信息转换成大写返回。
func handledata(conn net.Conn) {
	defer conn.Close()
	//fmt.Println("服务器端等待客户端消息")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		s := string(buf[:n])
		if s == "exit\n" || "exit\r\n" == s || s == "exit" {
			fmt.Println(conn.RemoteAddr(), "客户端要求退出")

			return
		}
		if n == 0 {
			fmt.Println("客户端已经关闭")
			return
		}
		if err != nil {
			fmt.Println(err)
			return

		}
		//s:=string(buf[:n])
		s = strings.ToUpper(s)
		fmt.Println(conn.RemoteAddr(), "客户端发来消息,已经回复：", s)
		conn.Write([]byte(s))
	}

}
