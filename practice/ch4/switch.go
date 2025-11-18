package main

import (
	"fmt"
	// "unicode/utf8"
	"math/rand"
)
func main () {
	// words := []string{"山", "sum", "微笑み", "人類学者","mountain", "タコの足とイカの足"}
	// for _, word := range words {
	// 	switch rc := utf8.RuneCountInString(word); rc{
	// 	case 1,2,3,4:
	// 		fmt.Printf("「%s」文字数は%dで短い単語だ。\n", word, rc)
	// 	case 5 :
	// 		bc := len(word)
	// 		fmt.Printf("「%s」の文字数は%dで、これでちょうどよい長さだ。ちなみにバイト数は%dだ。\n",word,rc,bc)
	// 	case 6,7,8,9:
	// 	default:
	// 		fmt.Printf("「%s」の文字数は%dで、とても長い！\n", word,rc)

	// 	}
	// }

	// wards2 := []string{"hi", "salutations","hello"}
	// for _, word := range wards2 {
	// 	switch wordLen := len(word); {
	// 	case wordLen < 5 :
	// 		fmt.Println(word, "は短い")
	// 	case wordLen > 10 :
	// 		fmt.Println(word, "は長すぎる")
	// 	default :
	// 		fmt.Println(word, "はちょうど良い長さの単語です")
	// 	}
	// }

	// // for文にloop（ラベル）を加えるとbreakでseitch文for文を抜け出す
	// loop:
	// for i := 0; i < 10; i++ {
	// 	switch {
	// 	case i%2 == 0 :
	// 		fmt.Println(i, ":偶数")
	// 	case i%3 == 0:
	// 		fmt.Println(i,":３で割り切れるが２で割り切れない")
	// 	case i%7 == 0:
	// 		fmt.Println(i,"：ループ終了…はしない！")
	// 		break loop
	// 	default:
	// 		fmt.Println(i, "：退屈な数")
	// 	}
	// }

	//goto文をの使用例
	// a := rand.Intn(10)
	// for a < 100 {
	// 	fmt.Println(a)
	// 	if a%5 == 0 {
	// 		goto done
	// 	}
	// 	a = a*2 + 1
	// }
	// fmt.Println("ループが通常終了したときに行う処理を実行")

	// done:
	// 	fmt.Println("ループが終わった時に必ず行う処理を実行")
	// 	fmt.Println(a)
}