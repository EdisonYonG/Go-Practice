package main

import "fmt"

// 定义全局变量 也可改成一次性声明
var n1 = 100
var n2 = 200
var name = "jack"

var (
	n3 = 300
	n4 = 400
	n5 = "marry"
)

func main() {

	// // 定义声明变量
	// var i int
	// i = 10
	// fmt.Println("i = ",i)
	// //第一种 指定类型变量 未定义则使用默认值
	// //int类型 默认值为零

	// //第二种 根据值自动判定变量类型（类型推导）
	// var num  = 10.11
	// fmt.Println("num=",num)

	// //第三种 省略var
	// // 下面的方式等价 var name string    name = "tom"
	// name := "tom"
	// fmt.Println("name = ",name)

	// var n1 int
	// var n3 int
	// var n2 int
	// 麻烦写法

	// var n1,n2,n3 int
	// fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)

	// 一次性声明多个变量方式2
	// var n1,name,n3 = 100,"tom",888
	// fmt.Println("n1=",n1,"name=",name,"n3=",n3)

	// 类型推导 第三种推导方式
	// n1,name,n3 := 100,"tom",888
	fmt.Println("n1=", n1, "name=", name, "n3=", n3)
	fmt.Println("n3=", n3, "n4=", n4, "n5=", n5)

	// 变量使用注意事项
	// 该区域的数据值可以在同一类型范围内不断变化
	// 变量在同一个作用域内不能重名
	// 变量=变量名+值+数据类型
	// golong变量如果没有赋初值，编译器会使用默认值

}
