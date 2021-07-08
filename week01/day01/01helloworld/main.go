package main
// 所有的.go文件都需要packeage，main包代表可以编译成可执行文件

import "fmt"
// import 代表倒导入，
// 函数外只能放标识符（变量、常量、函数、类型）的声明


// 声明变量
var name string;  // ""
var (
	age int; // 0
	address bool; // false
)

// 程序的入口函数
func main()  {
	fmt.Println("hello world");
	/**
	1、标识符：标识符是程序员定义的具有特殊含义的标识符，比如常量名、变量名、函数名等。Go语言的标识符由数字，字母，下划线组成，其中不能以数字开头。
	2、关键字和保留字不能作为变量名。
	3、变量规则
		1. go语言中的变量必须先声明再使用，
		2. 每个变量在声明时会有默认值，
		3. 局部变量声明之后必须被引用，不引用无法通过编译，不需要引用时可以用_做为变量名称。
		声明变量时必须要指定数据类型：
			var s1 string; int，bool
	4、go语言推荐驼峰命名法：如：studentName,
	*/
	name = "zs";
	fmt.Print(name);
	fmt.Printf("name:%s", name);
	fmt.Println(name);
	fmt.Print()

	// 5、声明变量的同时可以进行赋值
	var parents string = "parent";
	fmt.Println(parents,parents)
	// 6、类型推导：根据赋值的内容自动识别类型
	var company = "tencent";
	fmt.Println(company)
	// 7、简短变量声明，只能在函数内声明
	s3 := "hhh"
	fmt.Println(s3)
	// 8、匿名变量，匿名变量不占用命名空间，不分配内存。使用单下划线（_）声明
	var _ bool;
	// 9、同一个作用域中不能重复声明同名的变量
	// s3:="xixi"


}
