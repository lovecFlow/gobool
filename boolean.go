package main

import "fmt"

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// Выражение	Значение
// true && true = true
// true && false = false
// false && true = false
// false && false = false
// Выражение	Значение
// true || true = true
// true || false = true
// false || true = true
// false || false = false
// Выражение	Значение
// !true = false
// !false = true
//
