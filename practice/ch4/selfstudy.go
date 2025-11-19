package main

import (
	"fmt"
	// "math/rand"
)

func main() {
	// // 0以上100未満の乱数を合計100個、intのスライスに入れるforループを書け
	// numbers := []int{}
	// //rand(0)はパニックになるため初期値は「1」
	// for i := 1; i <= 100; i++ {
	// 	n := rand.Intn(i)
	// 	numbers = append(numbers, n)
	// }
	// // numbersは100の整数が入った配列
	// fmt.Println("長さ：",len(numbers),"中身：",numbers )

	// for _, num := range numbers{
	// 	if num %2 == 0 && num %3 == 0 {
	// 		fmt.Println("Six!")
	// 	} else if num %2 == 0 {
	// 		fmt.Println("Tow!")
	// 	} else if num %3 == 0 {
	// 		fmt.Println("Three!")
	// 	} else {
	// 		fmt.Println("Never mind")
	// 	}
	// }

	var total int
	for i := 0; i < 10; i++ {
		total := i + total
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Println(total)

	// 上記のようにしてしまうと「0」が出力される
	// 理由：変数のシャドウイング（shadowing）が発生している
	//
	// 正しい説明：
	// 1. スコープ内からスコープ外の変数にはアクセスできる（逆ではない）
	// 2. スコープ外でvar total intで宣言し、スコープ内でtotal := i + totalとしている
	// 3. := は短縮変数宣言で、新しい変数を宣言する（代入ではない）
	// 4. これはスコープ内で「var total int; total = i + total」をしているのと同義
	// 5. 外側のtotalと内側のtotalは別の変数（別メモリに格納）
	// 6. ループ内のtotalは外側のtotalの値（0）を参照して計算するが、
	//    外側のtotal自体は変更されない
	//
	// 解決策：total := i + total ではなく total = i + total（代入）を使用する

}
