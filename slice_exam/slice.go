package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slice example")

	// 리터럴 초기화 ( 배열과 같음 )
	var x = []int{1, 2, 3, 4, 5}

	fmt.Println(x)

	// 비교
	//var x2 = []int{1, 2, 3, 4, 5}
	//fmt.Println(x == x2) // ./slice.go:18:14: invalid operation: x == x2 (slice can only be compared to nil)

	var y = []int{3: 5, 10: 6}

	fmt.Println(y)

	var empty []int

	fmt.Println(empty)
	fmt.Println(empty == nil) // slice는 오직 nil 로만 비교 가능

	// append ( call by value )
	var a []int
	fmt.Println("a length:", len(a), a)

	a = append(a, 5)
	fmt.Println("a length:", len(a), a)

	a = append(a, 1, 2, 3)
	fmt.Println("a length:", len(a), a)

	// append(a, 4, 10) // append(a, 4, 10) (value of type []int) is not used - 반환된 값을 할당하지 않으면 오류 발생.

	// Capacity
	var str = []string{"a", "b", "c"}
	fmt.Println(str, len(str), cap(str))

	str = append(str, "d")
	fmt.Println(str, len(str), cap(str))

	str = append(str, "e", "f")
	fmt.Println(str, len(str), cap(str))

	// make
	var top5_fruits = make([]string, 5)
	fmt.Println(top5_fruits, len(top5_fruits), cap(top5_fruits))

	top5_fruits[0] = "apple"
	top5_fruits[1] = "banana"

	fmt.Println(top5_fruits, len(top5_fruits), cap(top5_fruits))

	// top5_fruits = make([]string, 5, 10)

	// fmt.Println(top5_fruits, len(top5_fruits), cap(top5_fruits))

	var rank = make([]int, 5, 10)
	fmt.Println(rank, len(rank), cap(rank))

	// Slicing
	var t = []int{1, 2, 3, 4, 5}
	var t1 = t[:3]
	var t2 = t[:]

	fmt.Println(t)
	fmt.Println(t1)
	fmt.Println(t2)

	t[0] = 10
	fmt.Println(t2)

	// Copy
	source := []string{"apple", "banana", "kiwi"}
	target := make([]string, 3)

	num := copy(target, source)

	fmt.Println(target, num)

	// length 크기에 따라 복사
	target2 := make([]string, 2)

	num = copy(target2, source)

	fmt.Println(target2, num)
}
