package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"strconv"
	_ "unsafe"
)

func main() {
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	// 这个函数会返回两个值 (value bool,err error)
	// 因为我只想获取value bool 就用下划线忽略err
	fmt.Printf("b type %T  b = %v\n", b, b)

	var str2 string = "1234590"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 = type %T n1 = %v\n", n1, n1)
	fmt.Printf("n2 = type %T n2 = %v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 = type %T f1 = %v\n", f1, f1)
	// string 转基本数据类型的注意事项
	// 在将 String 类型转成 基本数据类型时，
	// 要确保 String 类型能够转成有效的数据，
	// 比如 我们可以把 "123" , 转成一个整数，
	// 但是不能把 "hello" 转成一个整数，如果这样做，
	// Golang 直接将其转成 0 ， 其它类型也是一样的道理.
	// float => 0 bool => false
}
