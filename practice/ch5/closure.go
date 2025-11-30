package main

import (
	"fmt"
)


func main () {
	a := 20
	f := func() { //funcからが関数の定義。定義を変数fに代入
		fmt.Println(a) // 注意：外側で定義された変数にアクセス
		a := 30 //新たに変数aを宣言。main()のaはシャドーイングされる
		fmt. Println(a)//30
	} //ここまで関数の定義
	f()
	fmt.Println(a)// 20
	main2()
}

func makeMult(base int)func (int) int {
	return func (factor int) int {
		return base * factor
	}
}

func main2 () {
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	//関数自体を出力しようとするとメモリアドレスが返ってくる
	fmt.Println(threeBase, threeBase)
	for i := 0; i <= 5; i++ {
		fmt.Print(i, ": ", twoBase(i), ", ", threeBase(i), "\n")
	}
}
