package main

// byte和rune类型
// go语言为了处理非ascii外的类型，定义了一种新的类型：rune类型
// rune类型代表一个utf8的字符
import "fmt"

func main() {
	// 字符串修改
	s2 := "cloud"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
}
