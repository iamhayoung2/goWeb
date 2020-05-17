package main

import "fmt"

func main() {

	sum, diff := SumAndDiff(3, 2)
	fmt.Println(sum, diff)

	r := sum2(1, 2, 3, 4, 5)
	fmt.Println(r)

	n := []int{1, 2, 3, 4, 5} // 슬라이스 값 넘기기
	k := sum2(n...)
	fmt.Println(k)

	//함수를 변수에 저장하기
	var sumFunc func(a int, b int) int = sum1
	sumFunc2 := sum1

	fmt.Println(sumFunc(1, 2))
	fmt.Println(sumFunc2(1, 2))

	// 함수를 슬라이스에 저장하기
	f := []func(int, int) int{sum1, diff1}
	fmt.Println(f[0](1, 2), f[1](1, 2))

}

func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

func SumAndDiff2(a int, b int) (c int, d int) {
	c = a + b
	d = a - b
	return c, d
}

//가변인자 사용
func sum2(n ...int) int { // int형 가변인자를 받는 함수 정의
	total := 0
	for _, value := range n { // rnage로 가변인자의 모든 값을 꺼낸다.
		total += value
	}
	return total

}

//함수 저장하기
func sum1(a int, b int) int {
	return a + b
}

func diff1(a int, b int) int {
	return a - b
}

func nonNameFunc() {
	//익명함수
	//익명함수를 정의한 뒤 바로 호출한다.
	//익명함수는 함수를 정의한 뒤 () 괄호를 사용하여 바로 함수를 호출한다.
	// 코드 양을 줄이며, 뒤에서 다룰 클로저, 지연호출, 고루틴에서 사용

	func() {
		fmt.Println("함수에 이름없는 익명함수")
	}()

	func(s string) {
		fmt.Println(s)
	}("찍어서 바로 호출")

	r := func(a int, b int) int { // 익명 함수를 정의한 뒤
		return a + b
	}(1, 2) // 바로 호출하여 리턴값을 변수 r에 저장한다.

	fmt.Println(r) // 3

}
