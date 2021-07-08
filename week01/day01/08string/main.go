package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "武汉"
	fmt.Printf("%T\n", s1)
	b1 := 'h'
	fmt.Printf("%T\n", b1) // 	int32
	fmt.Printf("%d\n", b1) //		104
	path := "/Users/liuhuan/workspace/GoSpace/src/github.com/liuhuanxg/GoProjects/week01/day01/08string"
	fmt.Print(path)
	s2 := `海阔凭鱼跃,
		天高任鸟飞
	`
	fmt.Print(s2)

	//  字符串拼接
	name := "cloud"
	word := "云"
	ss := name + word
	fmt.Print(ss)
	// 占位符拼接
	s3 := fmt.Sprintf(name, word)
	fmt.Println(s3)
	paths := strings.Split(path, "/") // 分割
	fmt.Println(paths)
	fmt.Println(strings.Contains(s3, "cloud"))  // 包含
	fmt.Println(strings.HasPrefix(s3, "cloud")) // 是否以。。开头
	fmt.Println(strings.HasSuffix(s3, "cloud")) // 是否以//结尾
	s4 := "hello"
	fmt.Println(strings.Index(s4, "l"))		// 第一个下标
	fmt.Println(strings.LastIndex(s4, "l"))		// 最后一个的下标
	fmt.Println(strings.Join(paths,"+"))		// 字符串拼接
	fmt.Println(len(name))		// 字符串拼接
}
