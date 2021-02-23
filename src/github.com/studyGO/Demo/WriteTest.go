package main

/*
		Writer接口
		type Writer interface {
	    	Write(p []byte) (n int, err error)
		}
*/
import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println("创建文件错误", err)
		return
	}
	// defer用来声明一个延迟函数 在return执行之前执行 暂时理解为java中的finally 大多数时候用来释放资源
	defer file.Close()

	for i := 0; i < 5; i++ {
		file.WriteString("hello\n")
		file.Write([]byte("world\n"))
	}
}
