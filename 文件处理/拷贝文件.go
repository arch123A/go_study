package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ori := "/home/yu/go/src/test1/a.txt"
	des := "/home/yu/go/src/test2/a_bak.txt"
	if cpfile2(ori, des) {
		fmt.Println("successful")
	} else {
		fmt.Println("failed")
	}

}

//拷贝文件，如果成功返回true,否则返回false
func cpfile2(origin, destination string) bool {
	var finish bool = false
	file1, err1 := os.Open(origin)
	if err1 != nil {
		fmt.Println(err1)
		return finish
	}
	defer file1.Close()
	file2, err2 := os.Create(destination)
	if err2 != nil {
		fmt.Println(err2)
		return finish

	}
	defer file2.Close()
	buffer := make([]byte, 1024)
	for {
		n, err3 := file1.Read(buffer)
		if err3 == io.EOF {
			fmt.Println("finished!")
			finish = true
			break
		} else if err3 != nil {
			fmt.Println(err3)
			break
		}
		//n2,_:=file2.Seek(0,2)
		//file2.WriteAt(buffer[:n],n2)
		file2.Write(buffer[:n])

	}
	return finish
}
