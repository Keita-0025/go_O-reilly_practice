package main

import(
	"fmt"
)

func main() {
	// // inputの要素を逆順に出力
	// input := []int{1, 2, 3, 4, 5}
	// revers := make([]int, 0,len(input))
	// for i := len(input) -1; i >= 0; i-- {
	// 	revers = append(revers, input[i])
	// }
	// fmt.Println(revers)

	// //inputの要素の合計
	// input2 := []int{1, 2, 3, 4, 5}
	// sum := 0
	// for i := 0; i < len(input2); i ++ {
	// 	sum += input[i]
	// }
	// fmt.Println(sum)

	// // 配列から入力された番号のインデックスの要素を削除
	// idx := 0
	// fmt.Println("番号を入力してください")
	// fmt.Scan(&idx)
	// s := []int{10, 20, 30, 40, 50}

	// if idx < 0 || idx >= len(s) {
	// 	fmt.Println("invalid index")
	// }
	// s = append(s[:idx], s[idx+1:]...)
	// fmt.Println(s)

	// //整数の配列の重複をなくす
	// s := []int{3, 1, 4, 1, 5, 9, 3, 5}
	// seen := make(map[int]bool)
	// unique := make([]int, 0, len(s))
	// for _,v := range s {
	// 	if !seen[v] {
	// 		seen[v] = true
	// 		unique = append(unique, v)
	// 	}
	// }
	// fmt.Println(unique)

	// // 別解structを使う（超軽量はSet）
	// s := []int{3, 1, 4, 1, 5, 9, 3, 5}
	// seen := make(map[int]struct{})
	// unique := make([]int, 0, len(s))

	// for _,v := range s {
	// 	if _, ok := seen[v]; !ok {
	// 		seen[v] = struct{}{}
	// 		unique = append(unique, v)
	// 	}
	// }
	// fmt.Println(unique)

	// スライス s := make([]int, 0, 3) に順に append を行っていき、len と cap の変化を出力するプログラムを書け。出力は各 append 後の len と cap のペア。
	// 入力（操作系列）
	// 初期：s := make([]int, 0, 3)
	// append 1 → s = append(s, 1)
	// append 2 → s = append(s, 2)
	// append 3 → s = append(s, 3)
	// append 4 → s = append(s, 4)
	// append 5 → s = append(s, 5)
	// 出力（期待値）（len,cap の配列）
	// s := make([]int, 0, 3)
	// history := [][2]int {}
	// for i := 1; i < 5; i++ {
	// 	s = append(s, i)
	// 	history = append(history, [2]int{len(s),cap(s)})
	// }
	// fmt.Println(s, history)

	//別解
	// s2 := make([]int, 0, 3)
	// type SlicesInfo struct {Len, Cap int}
	// history := []SlicesInfo{}
	// for i := 1; i < 5; i++ {
	// 	s2 = append(s2, i)
	// 	history = append(
	// 		history,
	// 		SlicesInfo{
	// 			Len: len(s2),
	// 			Cap: cap(s2),
	// 		})
	// }
	// fmt.Println(history)

	// //別解2
	// s3 := make([]int, 0, 3)
	// history := []struct {
	// 	Len int
	// 	Cap int
	// }{}
	// for i := 1; i < 5; i++ {
	// 	s3 = append(s3, i)
	// 	history = append(history, struct{
	// 		Len int
	// 		Cap int
	// 		}{
	// 			Len: len(s3),
	// 			Cap: cap(s3),
	// 		})
	// }
	// fmt.Println(history)

	// スライスの部分操作で元配列が変わるか（上級）
	// 仕様：与えられた配列 a := [5]int{1,2,3,4,5} からスライス s := a[1:4] を作り、s[0] = 99 と代入した後の配列 a とスライス s を答えよ（つまり underlying array の共有を問う）

	// a := [5]int{1,2,3,4,5}
	// s := a[1:4] // 出力{2,3,4} 配列インデックス１から３まで
	// fmt.Println(s)
	// s[0] = 99 //出力{99,3,4}
	// fmt.Println(s,a)//出力{99,3,4}, {1,99,3,4,5}
	//スライスは配列を参照したビューである
	//スライスは配列の実態を共有する
	//ｓはただaの配列の1~3を参照しているだけ、aの裏にsがいるイメージ

	// 配列（実体）
	// [ 1   2   3   4   5 ]
	// 　     ↑--- s ---↑

	// スライス（Go内部の構造体）
	// {
	// 	ptr → 2 の場所（= a[1]）
	// 	len → 3
	// 	cap → 4
	// }

	// 2次元スライスのフラット化（上級）
	// 仕様：[][]int を受け取り全要素を1次元スライスに連結して返す Flatten([][]int) []int を実装せよ。内部で make を使って一度に正しい長さのスライスを確保すること（効率性チェック）。
	// 入力
	// m := [][]int{
	//  {1, 2, 3},
	//  {},
	//  {4, 5},
	// {6},
	// }

	// var flatten []int

	// for _, arrVal := range m {
	// 	flatten = append(flatten, arrVal...)
	// }
	// fmt.Println(flatten)

	//別解
	// total := 0

	// for _, arrVal := range m {
	// 	total += len(arrVal) //totalにインクリメント（＋）する
	// }
	// flatten := make([]int, 0, total)

	// for _, arrVal := range m {
	// 	flatten = append(flatten, arrVal...)
	// }
	// fmt.Println(flatten)

	// result := flatten(m)
	// fmt.Println(result)

	// スライスのコピーと独立性（上級）
	// 仕様：s1 := []int{1,2,3} を s2 := s1（直接代入）と s3 := make([]int, len(s1)); copy(s3, s1) で複製した後、s1[0] = 10 を行った場合の s1, s2, s3 の値を答えよ。目的は「代入による参照共有」と「copy による独立コピー」の違いを理解すること

	// s1 := []int{1,2,3}
	// fmt.Println(s1) //[1,2,3]
	// s2 := s1
	// fmt.Println(s1)//[1,2,3]
	// fmt.Println(s2)//[1,2,3]
	// s3 := make([]int, len(s1))
	// fmt.Println(s1)//[1,2,3]
	// fmt.Println(s2)//[1,2,3]
	// fmt.Println(s3)//[]
	// copy(s3, s1)
	// s1[0] = 10
	// fmt.Println(s1)//[10,2,3]
	// fmt.Println(s2)//[10,2,3]
	// fmt.Println(s3)//[1,2,3]
	//copyメソッドを使うと深いコピーになり独立した値となる（全く別物）
	//s2 := s1のようにするとただ単に参照しているだけになる（結局s2はs1のスライスを参照しているだけ）

	// スライスを使った窓（スライディングウィンドウ）和（応用）
	// 長さ n の整数スライス a と窓幅 k が与えられる。各連続部分列の合計（窓の和）をスライスで返す SlidingSums(a []int, k int) []int を実装せよ。k > len(a) のときは空スライスを返す

	a := []int{1, 8, 3, 4, 7, 6, 6, 6}
	k := 4
	// SlidingSum := func (a []int, k int)[]int{
	// 	if k <= 0 || k > len(a) {
	// 		fmt.Println("2つ目の引数には1つ目の引数の配列の長さより小さい数値を入れてください")
	// 		return []int{}
	// 	}
	// 	fmt.Println("OK,窓は",len(a)-k+1,"個できますね")
	// 	result := make([]int, 0, len(a)-k+1)
	// 	sum := 0
	// 	for i := 0; i < k; i++ {
	// 		sum += a[i]
	// 	}
	// 	result = append(result, sum)
	// 	for i := 1; i < len(a) -k+1 ; i++ {
	// 		sum = sum - a[i-1] + a[i+k-1]
	// 		result = append(result, sum)
	// 	}
	// 	return result
	// }

	// sums := SlidingSum(a, k)
	// fmt.Print(sums)
	//SlidingSumを使って平均を求める
	// avgs := make([]float64, 0, len(sums))
	// for i := 0; i < len(sums); i++ {
	// 	avg := float64(sums[i]) / float64(k)
	// 	avgs = append(avgs, avg)
	// }
	// fmt.Println(avgs)

	// SlidingAge := func (a []int, k int) [] float64{
	// 	if k <= 0 || k > len(a) {
	// 		return []float64{}
	// 	}
	// 	result := make([]float64, 0, len(a)-k+1)
	// 	sum := 0
	// 	for i := 0; i < k; i++ {
	// 		sum += a[i]
	// 	}
	// 	age := float64(sum) / float64(k)
	// 	result = append(result, age)
	// 	for i := 1; i < len(a)-k+1; i++ {
	// 		sum = sum - a[i-1] + a[i+k-1]
	// 		age = float64(sum) / float64(k)
	// 		result = append(result, age)
	// 	}
	// 	return result
	// }
	// avg := SlidingAge(a, k)
	// fmt.Println(avg)

	// SlidingMax := func(a []int, k int) []int {
	// 	if k <= 0 || k > len(a) {
	// 		return []int{}
	// 	}
	// 	result := make([]int, 0, len(a)-k+1)
	// 	max := 0
	// 	for i := 0; i < len(a)-k+1; i++ {
	// 		max = a[i]
	// 		for j := i + 1; j < i+k; j++ {
	// 			if a[j] > max {
	// 				max = a[j]
	// 			}
	// 		}
	// 		result = append(result, max)
	// 		fmt.Println(result)
	// 	}
	// 	return result
	// }
	// SlidingMax (a, k)
	
	// アルゴリズム
	// SlidingMax := func(a []int, k int) []int {
	// 	if k <= 0 || k > len(a) {
	// 		return []int{}
	// 	}
	// 	result := make([]int, 0, len(a)-k+1)
	// 	dq := make([]int, 0, len(a))
	// 	for i := 0; i < len(a); i++ { //len(a)
	// 		if len(dq) > 0 && dq[0] <= i-k {
	// 			dq = dq[1:]
	// 		}
	// 		for len(dq) > 0 && a[dq[len(dq)-1]]<=a[i] {
	// 			dq = dq [:len(dq)-1]
	// 		}
	// 		dq = append(dq, i)
	// 		maxVal := a[dq[0]]
	// 		if i >= k-1 {
	// 			result = append(result, maxVal)
	// 		}
	// 		fmt.Println("i", dq, "値", result)
	// 	}
	// 	return result
	// }
	// max := SlidingMax(a, k)
	// fmt.Println(max)

}

//関数の作成
// func flatten (towD [][]int) []int {
// 	total := 0
// 	for _, arrVal := range towD {
// 		total += len(arrVal)
// 	}
// 	result := make([] int, 0, total)
// 	for _, arrVal := range towD {
// 		result = append(result, arrVal...)
// 	}
// 	return result
// }
