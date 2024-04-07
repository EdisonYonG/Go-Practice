package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	_ "unsafe"
)

// 演示golong中基本数据类型的转换
// Golang 和 java / c 不同，
// Go 在不同类型的变量之间赋值时需要显式转换。强转
// 也就是说 Golang 中数据类型不能自动转换。

func main() {
	var i int32 = 100
	// 希望将i转为float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	fmt.Printf("n2 = %v,n1 = %v i = %v\n", n2, n1, i)

	// 在转换中，比如将 int64	转成 int8 【-128---127】 ，
	// 编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样。
	// 因此在转换时，需要考虑范围.

	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Printf("num1 = %v,num2 = %v", num1, num2)
}
