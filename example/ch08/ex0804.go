package main

import (
	"fmt"
)

// Map はマッピング関数fを使って[]T1を[]T2に変換（「マップ」）する //liststart1
// 2つの型パラメータT1、T2をとる（T1とT2は違ってもよい）
// 任意の型のスライスを扱える
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s)) // 変換結果を記憶するスライス
	for i, v := range s {
		r[i] = f(v)  // sの各要素を関数fで変換し結果をrに記憶
	}
	return r  // 変換結果が記憶されたスライスを返す
}

// Reduceは[]T1を、関数を用いて1個のT2の値に「縮約（ひとつの値で代表）」する
// （下のように呼び出された場合は、合計になる）
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filterは指定された関数を使ってスライスの要素をフィルタリングする
// この場合は関数fがtrueを返す要素だけからなるスライスを返す
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
} //listend1


func main() {
	words := []string{"One", "Potato", "Two", "Potato", "Three"} //liststart2
	filtered := Filter(words, func(s string) bool {
		return s != "Potato" // "Potato"でないものはtrue
	})
	fmt.Println(filtered) // [One Two Three]

	// Filterの結果のスライスfilteredの各要素に対して、lenを実行して
	// その結果を新しいスライスlengthsに代入
	lengths := Map(filtered, func(s string) int {
		return len(s) // 単語の文字数を返す（filteredの各要素を対象）
	})
	fmt.Println(lengths)  // [3 3 5]

	// Mapの結果のスライスlengthsの各要素の合計を計算
	// （このスライスを代表するものとして合計（sum）を計算する）
	sum := Reduce(lengths, 0, func(acc int, val int) int {
		return acc + val
	})
	fmt.Println(sum) // 11  //listend2

	// 別の例でMapを実行
	s := []int{10, 20, 30}
	s = Map(s, func(i int) int { // 自乗する
		return i*i
	})
	fmt.Println(s) // [100 400 900]
}
