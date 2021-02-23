package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookid  int
}

func main() {
	Book := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	fmt.Println(Book.title)
	fmt.Println(Book)
}
