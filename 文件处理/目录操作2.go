package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//打开文件判断文件是否t文件 ...
	f, err := os.OpenFile("/home/yu/go/src/test1/", os.O_RDONLY, os.ModeDir)
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

				fmt.Println(name)
			}
		}
	}

}
