package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	XMLName     xml.Name `xml:"student"` // 只有变量名字叫XMLName才能解析出来（结构体中标签的名字必须取名为XMLName，否则解析不出来。这是golang文档中的解释）
	StudentName string   `xml:"studentName"`
	StudentId   string   `xml:"studentId"`
}

type Students struct {
	XMLName     xml.Name  `xml:"students"`
	Version     string    `xml:"version,attr"`
	Stustruct   []Student `xml:"student"`
	Description string    `xml:",innerxml"`
}

func main() {
	//只读方式打开一个名称为 xxx 的文件
	file, err := os.Open("./student.xml")

	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	//调用ioutil包下的ReadAll函数（从 r 读取，直到出错或 EOF 并返回它读取的数据。）
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件流失败", err)
		return
	}
	defer file.Close()

	//解析XML文件
	stus := Students{}
	//Unmarshal函数解析XML编码的数据并将结果存入v指向的值。stus只能指向结构体、切片或者和字符串。良好格式化的数据如果不能存入stus，会被丢弃。
	err = xml.Unmarshal(data, &stus)
	if err != nil {
		fmt.Println("格式化XML数据失败", err)
		return
	}

	//输出XML中的数据
	// fmt.Println(string(data))
	fmt.Println("Description: " + stus.Description)
	fmt.Printf("XMLName: %v\n", stus.XMLName)
	fmt.Println("XMLVersion: ", stus.Version)

	for index, data := range stus.Stustruct {
		fmt.Printf("第%d个名字: %v\n", index+1, data.StudentName)
		fmt.Printf("第%d个ID: %v\n", index+1, data.StudentId)
	}
}
