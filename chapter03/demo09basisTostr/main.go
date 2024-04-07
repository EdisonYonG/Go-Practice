package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"strconv"
	_ "unsafe"
)

// golong基本数据类型练习转化为string使用
func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string

	// 第一种 fmt.Sprintf方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str = %q\n", str, str)

	// 第二种方法strconv包
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str = %q\n", str, str)
	// f = 格式 10 = 表示小数点保留十位 64 = 表示这个小数是float64

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str = %q\n", str, str)

	// strconv中有一个函数Itoa
	var num5 int = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str = %q", str, str)

}
