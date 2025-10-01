package main

import "fmt"

const x int64 = 10

const (
	key = "key"
	name = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)

	// x = x + 1 // エラー
	// y = "bye" // エラー

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// const w = x + y // エラー
	// fmt.Println(w)
}