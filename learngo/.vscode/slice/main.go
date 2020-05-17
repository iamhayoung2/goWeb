package main

import "fmt"

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
	// 실제로 없어지는 것이 아니고, 메모리에서 가르키는 주소값이 바뀌는 것.
}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]

}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[4:8]

	fmt.Println(b)

	c := a[4:]
	fmt.Println(c)

	d := a[:5]
	fmt.Println(d)

	for i := 0; i < len(a); i++ {
		var back int
		a, back = RemoveBack(a)
		fmt.Println(a)
		fmt.Println(back)
	}

	for i := 0; i < len(a); i++ {
		var back int
		a, back = RemoveFront(a)
		fmt.Println(a)
		fmt.Println(back)
	}

}
