package main

import "fmt"

func main() {
	// evenVals := []int{1,2,3,4,5,6,7,8,9,10} //偶数
	// for v := range evenVals {
	// 	if v%2 == 0 {
	// 		fmt.Println("偶数です:", v)
	// 	} else {
	// 		fmt.Println("奇数です:", v)
	// 	}
	// }

	// 省略する場合は_を使う
	//rangeは他の言語のイテレータと同じように使える
	// evenVals := []int{2,4,6,8,10,12} //偶数
	// for i,_ := range evenVals {
	// 	fmt.Println(i)
	// }

	// unipueNames := map[string]bool{"花子": true, "太郎": true, "洋子": true}
	// for k := range unipueNames {
	// 	fmt.Println(k)
	// }

	//_を使うときは値だけを使うorキーを捨てることを明示して意味があるときのみ使う

	// マップキーのイテレーションの順番は保証されない
	// m := map [string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }
	// for i := 0; i < 3; i ++ {
	// 	fmt.Println("ループ", i)
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }
	// go run forblock02.go && echo "---" && go run forblock02.go実行
	// 実行結果
	// ループ 0
	// b 2
	// c 3
	// a 1
	// ループ 1
	// c 3
	// a 1
	// b 2
	// ループ 2
	// a 1
	// b 2
	// c 3
	// ---
	// ループ 0
	// a 1
	// b 2
	// c 3
	// ループ 1
	// a 1
	// b 2
	// c 3
	// ループ 2

	//文字列に対するイテレーション
	// samples := []string{"hello", "apple_π!", "これは漢字文字列"}
	// for _, sample := range samples {
	// 	for i, r := range sample {
	// 		fmt.Println(i, r, string(r))
	// 	}
	// }

	//値を変更してもソースは変更されない
	// evenVals := []int{2,4,6,8,10,12}
	// for _, v := range evenVals {
	// 	v *= 2
	// }
	// fmt.Println(evenVals)
	// func spulitStringと宣言できるのはどうファイルトップレベスの関数のみ
	// splitString := func () {
	// 	samples := []string{"hello", "apple_π!", "これは漢字文字列"}
	// outer:
	// 	for _, sample := range samples {
	// 		for i, r := range sample {
	// 			fmt.Println(i, r, string(r))
	// 			if r == 'l' || r == 'は' {
	// 				continue outer
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}
	// }
	// splitString()

	splitCenter := func() {
		evenVal := []int{1,2,3,4,5}
		for i, v := range evenVal {
			if i == 0 {
				continue
			}
			fmt.Println(i, v)
			if i == len(evenVal) - 2 {
				break
			}
			fmt.Println()
		}
	}
	splitCenter()

	splitCenter2 := func () {
		evenVal := [] int {1,2,3,4,5}
		for i := 1; i < len(evenVal) - 1; i++ {
			fmt.Println(i, evenVal[i])
		}
	}
	splitCenter2()

	// goにはdo...whileは無いが工夫することで記述できる
	// for {
	// 	// まずやりたい処理（最低1回は走る）
	
	// 	if 終了条件 {
	// 		break
	// 	}
	// }
}