package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println("b = ", b)
	// bool类型占用存储空间是一个字节
	fmt.Println("b占用的空间 = ", unsafe.Sizeof(b))
	// b = 1 不能强转 只能取true和false
}
