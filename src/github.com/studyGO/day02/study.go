package main

import "fmt"

const s2 = 10000000000000000

func main() {
	// fmt.Println(s2)
	// printA()
	// i := returnA(10, 23)
	// fmt.Println(i)
	a := 10
	var prt *int
	var prt2 *float32
	prt = &a

	fmt.Printf("变量的地址: %x\n", prt)
	fmt.Printf("变量的地址: %x\n", *prt2)
}

func printA() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	strings := []string{"google", "runoob"}
	ints := []int{1, 3443, 4444}
	for _, s := range strings {
		fmt.Println(s)
	}
	fmt.Println(strings)
	fmt.Println(ints)

	fmt.Println(s2)
}

func returnA(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
