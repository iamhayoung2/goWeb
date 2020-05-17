package main2

import (
	"fmt"
	"strings"
)

const name5 bool = false

// name4 := "test" 축약형은 func 함수 안에서만 가능하다.
func main() {

	totalLength, upperName := lenAndUpper2("nicot")
	fmt.Println(totalLength, upperName)

	repeatMe("nico", "ha", "young", "test")

	totalLength, upperName = lenAndUpper("nico")
	fmt.Println(totalLength, upperName)

	totalLength2, _ := lenAndUpper("nico") // _ 은 ignore의 의미임.
	fmt.Println(totalLength2)

	fmt.Println(multiply(2, 2))

	fmt.Println("Hello world")

	const name string = "nico"
	//name = "Lynn"

	fmt.Println(name)

	var name2 string = "nico"
	name3 := "nico3"
	fmt.Println(name3)
	name2 = "lynn"
	fmt.Println(name2)
	//something.SayHell()

}

func multiply(a, b int) int {
	return a * b

}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, uppercase string) {

	//defer fmt.Println("i'm done") // 해당 function이 실행되고 난 후에 코드가 실행된다.
	defer fmt.Println("defer에 의해 실행됨" + uppercase)
	length = len(name)
	uppercase = strings.ToUpper(name)
	defer fmt.Println("defer에 의해 실행됨2" + uppercase)

	return // naked retrun

}
func repeatMe(words ...string) {
	fmt.Println(words) // 배열 형태로 return 함

}
