package main

import (
	"fmt"
)

func main () {
	m := map[int]string{
		1:"一番目",
		2:"二番目",
	}
	modeMap(m)
	fmt.Println(m)

	s := []int{1,2,3}
	modeSlice(s)
	fmt.Println(s)

	main2()
}

type person struct {
	name string
	age int
}

func modifyFails(i int, s string, p person) {
	i *= 2
	s = "さようなら"
	p.name = "Bob"
}

func main2 () {
	p := person{}
	i := 2
	s := "こんにちは"
	fmt.Println(i, s, p)
	// 実行時される瞬間にコピーされるので元のpではない
	// modifyFailsをが書き換えているのは丸ごとコピーされた構造体のnameフィールドである
	modifyFails(i, s, p)
}

func modeMap (m map[int] string) {
	//以下の値で上書きされる
	m[2] = "こんにちは"
	m[3] = "さようなら"
	delete(m, 1)
	fmt.Println(m)
}

func modeSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	fmt.Println(s)
	//容量に余裕があれば同じ配列を利用する。容量が足りなければ、新たに配列を確保するため、元の配列とは別ものになる
	s = append(s, 10)
	fmt.Println(s)
}