//package关键字用来声明该文件属于哪个包  main包表示可以编译成一个可执行文件
package main

//导入语句
import "fmt"

var s1 string = "你好你好"
var (
	name  string = "name"
	age   string = "age"
	money string
)

const s2 bool = true

//函数外面只能放变量、常量、函数、类型的声明 不能放逻辑语句
//main函数 程序的入口函数
func main() {

	fmt.Print(s2)                 //打印后不会换行
	fmt.Println("Hello world!!!") //打印后换行
	fmt.Printf("字符串是:%s", s1)     //占位符打印
	fmt.Println(name + " " + age)

	//Go语言中 局部变量声明了必须使用 否则编译器会报错
	var len = 10
	fmt.Println(len)
	//简短变量声明 只能在函数里面用
	s3 := "哈哈哈"
	fmt.Println(s3)

	var s4 int8 = 9
	fmt.Println(s4)

	s5 := 's'
	fmt.Printf("%T", s5)
	fmt.Println()

	//-----------------------------------------------------
	var s6 bool = false
	if s6 {
		fmt.Println("是true")
	} else {
		fmt.Println("是false")
	}

}
