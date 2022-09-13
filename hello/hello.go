package main

import (
	"fmt"
	"math"

	"rsc.io/quote"

	"math/rand"
)

func add(x int, y int) int { // you can omit the type from all but last (For e.g. (x, y int))
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// This is known as a "Naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

const Message = "Hello, world"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	fmt.Println(quote.Opt())

	rand.Seed(32)
	fmt.Println(rand.Int31(), math.Sqrt(4), math.Pi, add(3, 5))

	fmt.Println(swap("World", "Hello"))

	fmt.Println(split(17))

	var i, python, java = true, false, "no!"
	fmt.Println(i, j, python, java)

	// you can use := instead of "var"
	// But you have to know that Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	apple, banana, kiwi := true, false, true
	fmt.Println(apple, banana, kiwi)
	fmt.Printf("Type %T value is: %v\n", apple, apple)

	// !Important
	var default_i int
	var default_f float64
	var default_b bool
	var default_s string
	fmt.Printf("%v %v %v %q\n", default_i, default_f, default_b, default_s)

	// const ( Constants cannot be declared using the := syntax. )
	fmt.Printf("%T is %v\n", Message, Message)

	fmt.Println(Big>>99, Small)
}
