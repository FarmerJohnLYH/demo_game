package main // 意思是

import "fmt"

var LALALA int = 1000

func main() {
	var num = 1e99 - 1
	_ = num
	fmt.Println(num)
	// 打印 num 的 数据类型
	fmt.Printf("%T\n", num)
}
