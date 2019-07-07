package main
import "fmt"


func r2()(x,y string){
	x= "aaaa"
	y="bbbb"
	return
}


func feedme(portion,eaten int) int{
	eaten=portion+eaten
	if eaten>=5{
		fmt.Println("I'm full")
		fmt.Printf("I have eaten %v", eaten)
		return eaten
	}
	fmt.Println("I have eaten %v",eaten)
	return  feedme(portion,eaten)
}

func main()  {
	feedme(1,0)

	//fmt.Println(r2())
}