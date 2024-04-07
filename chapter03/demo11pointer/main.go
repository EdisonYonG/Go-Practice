package main

// import "fmt"
// import "unsafe"

// 演示golong指针类型使用

import (
	"fmt"
	_ "strconv"
	_ "unsafe"
)

func main() {
	var i int = 10
	// i的地址是什么? 取地址符 &i
	fmt.Println("i的地址 = ", &i)

	var ptr *int = &i
	// ptr 是一个指针变量 类型是*int 本身的值是&i
	// 指针存储的是地址
	fmt.Println("ptr = ", ptr)
	// 指针本身也有个地址 ptr
	fmt.Println("ptr 本身地址为 ", &ptr)
	// 4)获取指针类型所指向的值，使用：*，
	// 比如：var ptr *int,  使用*ptr 获取 ptr 指向的值
	fmt.Println("ptr 指向的值为 ", *ptr)

}
