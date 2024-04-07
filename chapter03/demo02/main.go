package main

import "fmt"

// 变量使用的注意事项
// 该区域的数据值可以在同一类型范围内不断变化
// 变量在同一个作用域内不能重名
// 变量=变量名+值+数据类型
// golong变量如果没有赋初值，编译器会使用默认值
func main() {
	// var i int = 10
	// i = 20
	// i = 50
	// fmt.Println("i = ",i)

	// i := 99
	// 变量在同一个作用域内不能重名
	// int 小数 都默认为0 string 默认为空串

	// 演示go中+号的使用
	var i = 1
	var j = 2
	var r = i + j
	fmt.Println("r=", r)

	var str1 = "hello"
	var str2 = "world"
	var res = str1 + str2 //做拼接操作
	fmt.Println("res=", res)

}
