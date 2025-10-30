package main

import "fmt"

func main() {

	// i := 0
	// for ; i <10; i++ {
	// 	fmt.Println(i)
	// }

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		if 1%2 == 0 {
			fmt.Println("偶数です:", i)
			i++
		} else {
			fmt.Println("奇数です:", i)
			i += 2
		}
	}

	// 0
	// 奇数です: 0
	// 3
	// 奇数です: 3
	// 6
	// 奇数です: 6
	// 9
	// 奇数です: 9

	// i := 1
	// for i < 100 {
	// 	fmt.Println(i)
	// 	i = i * 2
	// }

	//通常の書き方
	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		if i%5 == 0 {
	// 			fmt.Println (i, "３でも５でも割り切れる")
	// 		} else {
	// 			fmt.Println(i, "３で割り切れる")
	// 		}
	// 	}  else if i%5 == 0 {
	// 		fmt.Println(i, "５で割り切れる")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	// イディオム的な書き方（continueを使う）
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println (i, "３でも５でも割り切れる")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i,"３で割り切れる")
			continue
		}
		if i%5 ==0 {
			fmt.Println(i, "５で割り切れる")
			continue
		}
		fmt.Println(i)
	}





}
