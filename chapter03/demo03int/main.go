package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

// 变量的数据类型
// 演示go中整数类型的使用
func main() {
	var i int = 1
	fmt.Println("i = ", i)

	var j int16 = -129
	fmt.Println("j = ", j)

	var n uint16 = 129
	fmt.Println("n = ", n)

	// 介绍如何查看某个变量的数据类型 %T
	// fmt.Printf()可以用于格式化输出
	var n1 = 100
	fmt.Printf("n1 的 类型 %T \n", n1)

	//怎么查看变量占用字节大小和数据类型（使用较多）
	var n2 int64 = 10
	fmt.Printf("n2 的 类型 %T n2占用的字节是 %d\n", n2, unsafe.Sizeof(n2))

	// 遵循保大不保小原则
	// 保证程序正确运行下 尽量使用小一点的存储空间
	// 5)bit: 计算机中的最小存储单位。byte:计算机中基本存储单元。
	// [二进制再详细说] 1byte = 8 bit
	var age byte = 150
	fmt.Println("age = ", age)
}
