package main

import "fmt"

type Person struct {
	name    string
	age     int8
	address string
	job     JobClass
}

/*
What is iota?
iota is an identifier that is used with constant and which can simplify constant definitions
that use auto-increment numbers.
The iota keyword represents integer constant starting from zero.
*/
type JobClass int

// It is not enumeration. Look at this issue: https://github.com/golang/go/issues/19814
const (
	_ JobClass = iota
	Profession
	Student
	Engineer
)

func main() {
	hong := Person{
		name:    "gildong",
		age:     30,
		address: "Seoul",
		job:     Profession,
	}

	fmt.Println(hong)

	hong.age = 20

	fmt.Println(hong.age)

	set_age(hong, 18)

	fmt.Println(hong)
}

func set_age(person Person, age int8) {
	person.age = age
}
