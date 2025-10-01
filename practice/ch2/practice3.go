package main

import "fmt"


func main () {
	b := byte(255)
	smallI := int32(2147483647)
	bigI := uint64(18446744073709551615)
	b ++
	smallI = smallI + 1
	bigI = bigI + 1
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}