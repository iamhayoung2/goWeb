package main3

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main(){
	var num1 uint8 = math.MaxUint8

	fmt.Println(num1)


	var s1 string ="안녕하세요"

	var s2 string = `안녕
	  helloworld~`

	  fmt.Println(s1)
	  fmt.Println(s2)

	  fmt.Println(len(s1))

	  fmt.Println(utf8.RuneCountInString(s1))

	  //문자열 연산하기

	  var s3 string = "한글"
	  var s4 string = "한글"
	  
	  fmt.Println(s3==s4)
	  fmt.Println(s1==s3)
	  fmt.Println(s1+s3)

	  var s5 string = "GoGo"
	  fmt.Printf("%c\n", s5[1])

	  //상수사용하기
	  const age int = 10
	  const name = "hayoung"
	  const age2, name2 = 11, "hayoung2"
	  const(
		  age3, age4 = 12, 13
		  name3, name4 = "h4", "h5"
	  )

	  fmt.Println(age, name, "__", age2, name2, "__")
	  fmt.Println(age4, age4, name3, name4)


	}