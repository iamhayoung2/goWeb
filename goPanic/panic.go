// defer 함수 : try finally 구문처럼. 가장 마지막에 실행되는 것
// 프로그램이 잘못되어 에러가 발생한 뒤 종료되는 상황 = 패닉
// panic(에러 메시지) 형태로, 런타임 에러 뿐만 아니라 사용자가 직접 에러 발생 - 문법적인 에러는 아니지만, 로직에 따라 에러 처리 하고 싶을 때
// recover = try catch 구문
// 패닉이 발생 시 프로그램이 바로 종료x 예외처리 시 사용
// 변수 := revoer()
// 패닉에서 사용한 에러 메시지를 받아온다.
package main

import "fmt"

func main(){

}

func f(){
	defer func(){
		s := recover()
		fmt.Println(s)
	}

	a := [...]int{1,2}

	for i:=0;i<5;i++{ // 배열 크기를 벗어난 접근
		fmt.Println(a[i])

	}
}

func main(){
	f()

	fmt.Println("hello world!")
	// 런타임 에러가 중간에 발생하지만 f() 에서 
	// recover 함수로 복귀되어서. 해당 부분은 실행됨
	
}