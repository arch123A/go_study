package main

import (
	"fmt"

)
import "io/ioutil"

func main() {



	file, err := ioutil.ReadFile("aa.txt")
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Printf("%s\n",file)
	fmt.Printf("%s",file)



}
