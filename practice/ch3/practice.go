package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var i int
	// var n int
	// fmt.Println("数字を入力してください")
	// fmt.Scan(&i, &n)
	// sum := i + n
	// fmt.Println(sum)

	// a, b := 2,1
	// a, b = b, a
	// fmt.Println(a, b)

	// x := 10
	// y := 5
	// x,y = x+1, y
	// sum := x + y
	// fmt.Println(sum)

	// var x int = 40
	// y := 2
	// sum := x + y
	// fmt.Println(sum)

	// c := 1
	// {
	// 	c := "inner"
	// 	fmt.Println(c)
	// }
	// fmt.Println(c)

	v := ""
	fmt.Println("値を入力してください")
	fmt.Scan(&v)

	n, err := strconv.Atoi(v)

	if err == nil {
		fmt.Println(n * 2)
	} else {
		fmt.Println("invalid input:", err)
	}

}
