package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{"Richard"}
	/// The only way to mutate a variable that you pass to a function is by passing a pointer.
	/// By default, the pass-by-value means that changes you make are on the copy you’re working on.
	/// Thus, they are not reflected in the calling function.
	p1 := rename(p)
	fmt.Println(p1, &p1.name)
	fmt.Println(p, &p.name)

	// If you want to mutate person, you could just accept a pointer.
	rename_p(&p)
	fmt.Println(p, &p.name)
}

func rename(p person) person {
	p.name = "test"
	return p
}

func rename_p(p *person) {
	p.name = "test"
}
