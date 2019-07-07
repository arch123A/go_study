package main
import ("fmt"
	"strings"
)


func main()  {
	s := "oh I do like to be beside the seaside "
	//
	//fmt.Println("%q",s[0])
	//fmt.Printf("%q\n",s[0])
	//fmt.Printf("%b\n",2000)
	//fmt.Println(strings.ToUpper(s))
	//fmt.Println(strings.Index("surface","face2"))
	//s1 := "   \t sdsdsdsdsdsdd sddsds sddsds dsdssd sdsd   \t"
	//fmt.Println(strings.TrimSpace(s1))
	//fmt.Println("中国人")
	s1:=strings.ToUpper(s)
	s2 := strings.Replace(s, "seaside","bar",1)
	s3:=strings.Index(s, "the")
	fmt.Println(s1,s2,s3)
}