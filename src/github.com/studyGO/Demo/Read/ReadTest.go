package main

/*
		Reader接口
		type Reader interface {
   			 Read(p []byte) (n int, err error)
		}

*/
import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
	}

	//创建读取文件缓存区的数组(一次读取128个字节)
	var buffer [128]byte
	//文件内容字节数组
	var content []byte

	//无限循环（直至文件读取完成再跳出循环）
	for {
		n, err := file.Read(buffer[:])
		if err == io.EOF {
			//报错了就相当于 文件读取完毕
			break
		}
		if err != nil {
			fmt.Println("读取文件错误", err)
			return
		}
		defer file.Close()
		content = append(content, buffer[:n]...)
	}
	//输出读取的结果
	fmt.Println(string(content))
}
