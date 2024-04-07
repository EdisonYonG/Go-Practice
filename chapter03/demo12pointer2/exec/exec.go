package main

// import "fmt"
// import "unsafe"

// 演示指针练习

import (
	"fmt"
	_ "strconv"
	_ "unsafe"
)

func main() {
	var num int = 9
	fmt.Println("num address = ", &num)

	var ptr *int
	ptr = &num
	*ptr = 10
	//这里修改值 会导致num的变化
	fmt.Println("num = ", num)
}
