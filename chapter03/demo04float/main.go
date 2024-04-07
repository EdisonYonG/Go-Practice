package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
)

// 浮点数 = 符号位 + 指数位 + 尾数位
func main() {
	// var price float32 = 89.12
	// fmt.Println("price = ",price)

	// var num1 float32 = -0.0008912
	// var num2 float64 = -8687685.09
	// fmt.Println("num1 = ",num1,"num2 = ",num2)

	// // 尾数部分可能丢失 造成精度损失
	// var num3 float32 = -123.0000901
	// var num4 float64 = -123.0000901
	// fmt.Println("num3 = ",num3,"num4 = ",num4)

	// golong 浮点型数据类型默认为float64
	var num5 = 1.1
	fmt.Printf("num5的数据类型是 %T\n", num5)

	// 科学计数法
	num8 := 5.1234e2   //乘10的2次方
	num9 := 5.1234e2   //乘10的2次方
	num10 := 5.1234e-2 //除10的2次方
	fmt.Println("num8 = ", num8, "num9 = ", num9, "num10 = ", num10)
}
