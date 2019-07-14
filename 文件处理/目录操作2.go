package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//打开文件判断文件是否txt文件 ,如果是则拷贝到相应的文件

	path := "/home/yu/go/src/test1/"
	dest := "/home/yu/go/src/test2/"
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)

	}
	defer f.Close()
	fmt.Printf("%T\n", f)
	re, err := f.Readdir(-1)

	/*	Name() string       // base name of the file
		Size() int64        // length in bytes for regular files; system-dependent for others
		Mode() FileMode     // file mode bits 文件权限 -rw-r--r--
		ModTime() time.Time // modification time
		IsDir() bool        // abbreviation for Mode().IsDir()
		Sys() interface{}   // underlying data source (can return nil)*/

	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("%T\n", re)
	for _, value := range re {
		if !value.IsDir() {
			name := value.Name()
			if strings.HasSuffix(name, "txt") {
				ori := path + name
				b := cpfile(ori, dest+name)
				if b {
					fmt.Println(name, "拷贝成功")
				}

				fmt.Println(name)
			}
		}
	}

}

//拷贝文件，如果成功返回true,否则返回false
func cpfile(origin, destination string) bool {
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
			//fmt.Println("finished!")
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
