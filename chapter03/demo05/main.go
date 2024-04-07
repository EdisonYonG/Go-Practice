package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
)

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1 = ", c1)
	fmt.Println("c2 = ", c2)
	fmt.Printf("c2 = %c c1 = %c\n", c2, c1)

	// var c3 byte = '北'//overflow溢出
	// var c3 int = '北'//overflow溢出
	c3 := '北'
	fmt.Printf("c3 = %c c3对应码值为%d \n", c3, c3)

	// 1)如果我们保存的字符在 ASCII 表的,比如[0-1, a-z,A-Z..]
	// 直接可以保存到 byte
	// 2)如果我们保存的字符对应码值大于 255,
	// 这时我们可以考虑使用 int 类型保存
	// 3)如果我们需要安装字符的方式输出，
	// 这时我们需要格式化输出，即 fmt.Printf(“%c”, c1)

	// 按格式化输出要加%c
	c4 := 22269
	fmt.Printf("c4 = %c", c4)

	var n1 = 10 + 'a'
	fmt.Println("n1 = ", n1)

}
