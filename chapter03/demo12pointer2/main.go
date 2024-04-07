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
	var num int = 9
	fmt.Println("num的地址为:", &num)
	// var ptr *int = &num
	// fmt.Println("ptr的地址为：",ptr)
	// fmt.Println("ptr指向的值为：\n",*ptr)

	ptr := &num
	fmt.Println("ptr指向的值为:", *ptr)

}
