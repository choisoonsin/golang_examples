package main

import (
	"fmt"
	"strconv"
)

func main() {
	var fruits [3]string

	fmt.Println(fruits, len(fruits))

	// This is an efficient way
	// OKay you can checkout this benchmark at https://gist.github.com/evalphobia/caee1602969a640a4530
	for i := 0; i < len(fruits); i++ {
		fruits[i] = strconv.Itoa(i)
		fmt.Println(i, fruits[i])
	}

	for i := 0; i < len(fruits); i++ {
		fruits[i] = strconv.FormatInt(int64(i), 10)
		fmt.Println(i, fruits[i])
	}

	// You can also use fmt.Sprintf. but as it metioned above, you should use Itoa for performance
	for i := 0; i < len(fruits); i++ {
		fruits[i] = fmt.Sprint(i)
		fmt.Println(i, fruits[i])
	}
}
