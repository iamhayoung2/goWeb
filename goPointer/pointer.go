package main

import "fmt"

func main() {
	var a int
	var p *int

	a = 3
	p = &a

	fmt.Println(a)
	fmt.Println(&a, "__", p)
	fmt.Println(*p)
}

func Increase(x int) {

}