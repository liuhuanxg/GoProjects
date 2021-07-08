package main

import "fmt"
import "math"


func main(){

	var f1 = math.MaxFloat32
	fmt.Printf("%f\n", f1)
	var f2 = 1.234565
	fmt.Printf("%T\n", f2)  // float 64

	var f3 = float32(1.23456)
	fmt.Printf("%T\n", f3)  // float 32
	//f1 = f3 // float32 不能直接赋值给float64
}
