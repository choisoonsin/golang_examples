package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 배열
	var x [10]int // 제로값으로 초기화
	fmt.Println("x의 값은?", x)

	var fruits = [3]string{"apple", "banana", "kiwi"} // 배열 리터럴 사용
	fmt.Println("fruits의 값은?", fruits)

	var x1 = [3]int{1, 2, 3}
	var x2 = [...]int{1, 2, 3}
	fmt.Println(x1 == x2)

	fmt.Println(len(x1))

	x1[0] = 3
	x1[1] = 2
	x1[2] = 1
	fmt.Println(x1)

	var y = [...]int{1, 5: 6, 10: 1}
	fmt.Println(y)

	var yy = [...]int{10: 1}
	fmt.Println(yy)

	// Array Type 확인
	_type := reflect.TypeOf(yy)

	if _type.Kind() == reflect.Array {
		fmt.Println("type of yy is", _type.Elem())
	}
}
