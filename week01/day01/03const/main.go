package main

import "fmt"
/**
常量：定义之后不能修改，在程序运行期间不会改变的量
*/

// 声明单个常量
const pi=3.1415926

// 批量声明常量，如果某个常量之后的常量没有赋值，默认使用当前常量的值为后续变量赋值。
const (
	n1 = 100;
	n2
	n3
)
// iota在const中首次出现时默认置为0，之后const中每加一行常量，iota加1，
const (
	a1 = iota; // 0
	a2  // 1
	_   // 2
	a3  // 3
)

const (
	a4 = iota; // 0
	a5 = 100 // 1
	a6  = iota // 2
	a7  // 3
)
const (
	d1,d2 = iota+1, iota+2  // 1 2
	d3,d4 = iota+1, iota+2  // 2 3
)

// 3、定义数量级
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

func main()  {
	//fmt.Println(n1);
	//fmt.Println(n2);
	//fmt.Println(n3);
	//fmt.Println(a1);
	//fmt.Println(a2);
	//fmt.Println(a3);
	fmt.Println(a4);
	fmt.Println(a5);
	fmt.Println(a6);
	fmt.Println(d1);
	fmt.Println(d2);
	fmt.Println(d3);
	fmt.Println(d4);
	fmt.Println(KB, MB, GB, TB);

}
