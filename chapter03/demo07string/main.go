package main

// import "fmt"
// import "unsafe"

// 字符串就是一串固定长度的字符连接起来的字符序列。
// Go 的字符串是由单个字节连接起来的。

// 演示string类型使用

import (
	"fmt"
)

func main() {

	var address string = "北京长城 helloworld"
	// fmt.Println("地址是：",address)
	fmt.Println(address)

	// 2)字符串一旦赋值了，字符串就不能修改了：在 Go 中字符串是不可变的。
	// var str = "hello"
	// str[0] = 'a'  //这里不能修改str的内容

	// (2)反引号，以字符串的原生形式输出，包括换行和特殊字符，
	// 可以实现防止攻击、输出源代码等效果
	str2 := "abc\nabc"
	fmt.Println(str2)

	// str3 := `

	// var address string = "北京长城 helloworld"
	// // fmt.Println("地址是：",address)
	// fmt.Println(address)

	// // 2)字符串一旦赋值了，字符串就不能修改了：在 Go 中字符串是不可变的。
	// // var str = "hello"
	// // str[0] = 'a'  //这里不能修改str的内容

	// // (2)反引号，以字符串的原生形式输出，包括换行和特殊字符，
	// // 可以实现防止攻击、输出源代码等效果
	// str2 := "abc\nabc"
	// fmt.Println(str2)
	// `

	//反引号全部输出

	// 字符串拼接方式
	var str = "hello" + "world"
	str += "haha!"
	fmt.Println(str)
	// 很长的拼接操作可以分行写
	// 加号必须保留在上一行 若没有+ go自动加分号 导致程序报错

	var str4 = "hello" + "world" + "hello" + "world" +
		"hello" + "world" + "hello" + "world" + "hello" +
		"world" + "hello" + "world"
	fmt.Println(str4)

	var a int
	var b float32
	var c float64
	var isMarryied bool
	var name string
	fmt.Printf("a=%d,b=%f,c=%f,isMarryied=%v,name=%v", a, b, c, isMarryied, name)
	// %v = 按值的默认格式默认值输出  %f = 按标准计数法输出

}
