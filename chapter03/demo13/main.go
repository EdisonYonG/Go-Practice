// 1)由 26 个英文字母大小写，0-9  ，_ 组成
// 2)数字不可以开头。var num int //ok	var 3num int //error
// 3)Golang 中严格区分大小写。
// var num int
// var Num int
// 说明：在 golang 中，num 和 Num  是两个不同的变量
// 4)标识符不能包含空格。

// import "fmt"
// import "unsafe"

// 演示标识符命名规则练习
//
//	1.package进行包的声明  建议：包和所在文件夹同名
package main

import (
	"demo13/model"
	"fmt"
)

func main() {
	model.New()
	fmt.Println("main")
}

// package main
// import (
// 	"demo13_identifier/model"
// 	"fmt"
// 	_"unsafe"
// 	_"strconv"
// )

// func main()  {
// 	hello	// ok hello12 //ok
// 1hello // error ,不能以数字开头
// h-b	// error ,不能使用 -
// x h	// error, 不能含有空格

// h_4	// ok
// _ab	// ok
// int	// ok , 我们要求大家不要这样使用
// float32 // ok ,  我们要求大家不要这样使用
// _	// error Abc		// ok

// 包名：保持 package 的名字和目录保持一致，尽量采取有意义的包名，
// 简短，有意义，不要和标准库不要冲突 fmt

// 1)变量名、函数名、常量名：采用驼峰法举例：
// var stuName string = “tom”	形式: xxxYyyyyZzzz ...

// var goodPrice float32 = 1234.5
// 2)如果变量名、函数名、常量名首字母大写，则可以被其他的包访问；
// 如果首字母小写，则只能在本包中使用 ( 注：可以简单的理解成，
// 首字母大写是公开的，首字母小写是私有的) ,
// 在 golang 没有public , private 等关键字。

// hello pkg
// hello kkk
