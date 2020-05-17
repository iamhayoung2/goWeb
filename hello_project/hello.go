package main

import (
	"fmt"
	_ "time"
)

func main() {

	var x, y int = 30, 50
	_ = x // 사용하지 않는 변수

	fmt.Println(y)
	fmt.Println("hello world")

	z := " := 변수로 처리하기 "
	fmt.Println(z)

	var num1 int = 2
	var num2 = 3
	var num3 float32 = 0.1
	var a int

	fmt.Println(num1, num2, num3)

	for i := 0; i < 4; i++ {
		a = num1 + i
		fmt.Println(a)

	}

}
