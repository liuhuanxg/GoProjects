package main

import "fmt"

func main(){
	// 十进制
	var  a = 10;
	fmt.Printf("a:%d\n", a)
	// 八进制
	a2 := 077
	fmt.Printf("%d\n",a2)
	// 十六进制
	var a3 = 0xff;
	fmt.Printf("%d\n",a3)

	// 查看变量的类型
	fmt.Printf("%T\n", a3)
	// 声明int8类型的变量
	a4 := int8(9)
	fmt.Printf("%T\n", a4)
}


