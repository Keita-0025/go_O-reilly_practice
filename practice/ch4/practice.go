package main

import (
	"fmt"
)
func main() {
	var r float64
	fmt.Println("数字を入力してください")
	fmt.Scan(&r)
	const Pi = 3.14159265
	area := Pi * r * r
	fmt.Println(fmt.Sprintf("%.2f", area))

	var d int
	fmt.Println("数字を入力してください")
	fmt.Scan(&d)
	const (
		Sun = iota
		Mon
		Tue
		Wed
		Thu
		Fri
		Sat
	)

	switch d {
	case Sun:
		fmt.Println("Sun")
	case Mon:
		fmt.Println("Mon")
	case Tue:
		fmt.Println("Tue")
	case Wed:
		fmt.Println("Wed")
	case Thu:
		fmt.Println("Thr")
	case Fri:
		fmt.Println("Fri")
	case Sat:
		fmt.Println("Sat")
	default:
		fmt.Println("invalid")
	}
}
