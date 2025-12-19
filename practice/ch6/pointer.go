package main

import (
	"fmt"
)

func main() {
	x := 10
	pointerTox := &x
	fmt.Println(pointerTox)
	fmt.Println(*pointerTox)
	z := 5 + *pointerTox
	fmt.Println(z)
	var y *int
	fmt.Println(y == nil)
	// nilのポインタを デリファレンス（ポインタが参照するアドレスに保存されている値を返す）するとパニックになる
	// fmt.Println(*y)

	//varを使ってポインタ型の変数をを宣言する時は
	//そのポインタが指す領域に保存される値の型の前に「*」をつけて表す
	//すべての型について、ポインタ型を作ることができる
	w := 10
	var pointerTow *int
	pointerTow = &w
	fmt.Println(pointerTow)

	//newはポインタ型生成する
	//指定された型でゼロ値を値とするインスタンスへのポインタを返す
	var v = new(int) //vの参照先はintのゼロ値(0)が記憶される
	fmt.Println(v == nil)
	fmt.Println(v)//vのアドレスが表示される
	fmt.Println(*v == 0)
	fmt.Println(*v)

	type Foo struct{}
	u := &Foo{}
	var t string
	r := &t
	fmt.Println(u,r)

	type person struct {
		FirstName string
		MiddleName *string
		LastName string
	}
	p := person{
		FirstName: "Pet",
		// MiddleName: "Perry", //コンパイルError
		LastName: "Peterson",
	}
	fmt.Println(p)


	p2 := person {
		FirstName: "Pet",
		MiddleName: makePointer2("Perry"),
		LastName: "Prterson",
	}
	fmt.Println(p2)
	fmt.Println(*p2.MiddleName)
}

// 定数をポインタに変化するにはポインタに変換するにはヘルパー関数を使う
func makePointer[T any](t T) *T {
	return &t
}
func makePointer2(t string) *string{
	return &t
}