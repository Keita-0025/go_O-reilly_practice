package main

import (
	"fmt"
	"math/rand"
)


func main () {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("少し小さすぎます:", n)
	} else if n > 7 {
		fmt.Println("大きすぎます:", n)
	} else {
		fmt.Println("いい感じの数字です:", n)
	}

	if n := rand.Intn(10); n ==10 {
		fmt.Println("10です:", n)
	}else if n >= 5 {
		fmt.Println("半分以上です:", n)
	}else {
		fmt.Println("だめだこりゃ:", n)
	}
	// fmt.Println("nの値は", n, "です") //特殊なスコープ、ブロック外でnは使えない
}