package main

import "fmt"

func main() {
	fruitsPrice := map[string]int{}

	fruitsPrice["apple"] = 100
	fruitsPrice["kiwi"] = 90

	fmt.Println(fruitsPrice, len(fruitsPrice))

	v, ok := fruitsPrice["apple"]
	fmt.Println("apple", v, ok)

	v, ok = fruitsPrice["banana"]
	fmt.Println("banana", v, ok)

	fruitsPrice["banana"] = 200

	v, ok = fruitsPrice["banana"]
	fmt.Println("banana", v, ok)

	delete(fruitsPrice, "apple")

	v, ok = fruitsPrice["apple"]
	fmt.Println("apple", v, ok)

	// Set 만들기
	fruitSet := map[string]bool{}
	data := []string{"banana", "kiwi", "apple", "tomato", "watermelon", "melon", "banana", "tomato"}

	for _, v := range data {
		fruitSet[v] = true
	}

	fmt.Println(fruitSet, len(fruitSet), len(data))

}
