package main

//import "fmt"
import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(100 * time.Millisecond)
	var a []int
	fmt.Printf("len(a) : = %d\n", len(a))
	fmt.Printf("cap(a) : = %d\n", cap(a))

	b := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(b) : = %d\n", len(b))
	fmt.Printf("cap(b) : = %d\n", cap(b))

	c := make([]int, 5, 8)
	fmt.Printf("len(c) : = %d\n", len(c))
	fmt.Printf("cap(C) : = %d\n", cap(c))

	b = append(b, 8)
	fmt.Printf("len(b) : = %d\n", len(b))
	fmt.Printf("cap(b) : = %d\n", cap(b))

	e := []int{1, 2}
	f := append(e, 3)

	fmt.Printf("주소는? %p , %p \n", e, f) // e,f는 서로 다른 주소 값으로 나옴.

	for i := 0; i < len(e); i++ {
		fmt.Printf("%d", e[i])
	}

	fmt.Println()
	for i := 0; i < len(f); i++ {
		fmt.Printf("%d", f[i])
	}
	fmt.Println()
	fmt.Println(cap(e), " ", cap(f))

	aa := make([]int, 2, 8)
	aa[0] = 1
	aa[1] = 2
	bb := append(aa, 3)

	fmt.Println("aa와 bb")

	fmt.Printf("%p   %p", aa, bb)

	fmt.Println(cap(aa), " ", cap(bb))
	fmt.Println(aa, " ", bb)

	ab := make([]int, 2, 4)
	ab[0] = 1
	ab[1] = 2

	ac := make([]int, len(ab))

	for i := 0; i < len(ab); i++ {
		ac[i] = ab[i]

	}

	ac = append(ac, 3)

	fmt.Printf("%p   %p", ab, ac)

}
