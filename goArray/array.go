package main

import (
	"fmt"
	f "fmt"
)

func main() {
	//array()
	//array2()
	slice_ref()
}

func array() {

	var a [5]int // int형 길이 5 배열
	a[2] = 7
	f.Println(a)

	var b [5]int = [5]int{1, 2, 3, 4, 5}
	var c = [5]int{0, 1, 2, 3, 4}
	d := [5]int{1, 2, 3, 4, 5}

	_, _, _ = b, c, d

}

func array2() {
	b := [5]int{
		1,
		2,
		3,
		4,
		5, // 여러 줄로 나열 시 마지막 요소에도 콤마를 붙인다.
	}

	c := [...]int{1, 2, 3, 4} // 배열크기 ... 로 쓰면 자동 초기화

	_ = b

	f.Println(c[2])

	for i := 0; i < len(b); i++ {
		f.Println(c[i])
	}

}

func array3() {
	a := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	_ = a
}

func array_range() {
	a := [5]int{10, 11, 12, 13, 14}

	for i, value := range a {
		f.Println(i, value)
	}

	for value := range a {
		f.Println(value)
	}

	for _, value := range a {
		f.Println(value)
	}
}

func array_copy() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a // 배열을 대입하면 배열 전체가 복사됨.

	_, _ = a, b
}

func slice() { // 슬라이스는 배열과 같지만, 길이가 고정되어 있지 않아, 크기가 동적으로 늘어남.
	// 배열과 달리 레퍼런스 타입임.
	// [] 대괄호 안에 길이 지정 안함.
	// make 함수를 사용하여 공간을 할당해야 값을 넣을 수 있다.
	// make([]자료형, 길이) --> 길이 조정안하면 0 임.
	var a []int = make([]int, 5) // make 함수로 int형에 길이가 5인 슬라이스에 공간 할당, 초기값은 모두 0으로 셋팅됨
	var b = make([]int, 5)
	c := make([]int, 5)

	_, _, _ = a, b, c

	d := []int{1, 2, 3, 4, 5} // 슬라이스 생성과 동시에 초기화

	var e = make([]int, 5, 10) // 길이가 5이고 용량이 10
	// 미리 용량을 설정하면 요소가 추가될때마다 메모리를 새로 할당하지 않아. 성능상 이점. 단, 처음부터 메모리 공간을 많이 차지함.

	e = append(e, 50, 51)

	// 슬라이스에 값 추가하기
	d = append(d, 10, 11)

	f.Println(d)

	d = append(d, e...) // 슬라이스에 슬라이스를 붙일때는 e... 를 쓴다.

}

func slice_ref() {
	// 슬라이스는 레퍼런스 타입
	// 내장된 배열에 대한 포인터이므로, 슬라이스끼리 대입하면 값이 복사되지 않고 참조(레퍼러스)만 됨
	//cf. 배열과 비교
	a := [3]int{1, 2, 3}
	var b [3]int

	b = a
	b[2] = 9

	f.Println(a)
	f.Println(b)

	// 슬라이스
	c := []int{1, 2, 3}
	var d []int

	d = c
	d[3] = 9

	// 이럴경우, c[3] 도 9로 바뀜
	// 왜냐면 point라서 주소를 복사했기 때문. 값 복사 x

	// 슬라이스를 복사하려면 copy함수 사용
	//copy(복사될슬라이스, 원본슬라이스)

	e := []int{1, 2, 3, 4, 5}
	f := make([]int, 3)

	copy(f, e) //

	fmt.Println(e) // 1,2,3,4,5
	fmt.Println(f) // 1,2,3

	f[0] = 9 // f만 바뀌고, e는 안바뀜

	// 부분 슬라이스 - 참조타입

	g := e[0:5] // 끝 인데스는 +1 해야함. --> 1,2,3,4,5
	fmt.Println(g)

	h := e[1:4]   // 2,3,4
	i := e[1:4:5] // 용량 값
	_, _ = h, i

}
