package main

import "fmt"

func main() {
	//1.次の5つの要素を含む文字列スライス型のgreetingsという名の変数を定義する
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Mọi người mình là Kim"}

	//最初の2つの要素からなるスライス
	greetings1 := greetings[:2]
	fmt.Println(greetings1)

	//2~4番目の要素からなるスライス
	greetings2 := greetings[1:4]
	fmt.Println(greetings2)

	//4,5番目からなるスライス
	greetings3 := greetings[3:]
	fmt.Println(greetings3)

	//2.文字列に関して次の操作を行うプログラムを書け
	//"Hi 👩 and 👨"と定義
	//絵文字は4byteある
	message := "Hi 👩 and 👨"
	//4番目の文字を表示
	Three := message[12:]
	fmt.Println(len(message))
	fmt.Println(Three)

	//3.構造体に関して以下の操作を順番に行うプログラムを書け
	//firstName,lastName,及びidを持つ構造体を定義
	type Employee struct {
		firstName string
		lastName string
		id int
	}

	//この構造体の次の3種類のインスタンスを生成する
	//名前なしで構造体リテラルを使って初期化するもの
	keita := struct {
		firstName string
		lastName string
		id int
	}{
		firstName: "tanemori",
		lastName: "keita",
		id: 777,
	}
	fmt.Println(keita)

	//名前付きの構造体リテラルを使って初期化するもの
	mayu := Employee{
		firstName: "tanemori",
		lastName: "mayu",
		id: 778,
	}
	fmt.Println(mayu)

	//varを使って初期化するもの
	var yua struct {
		firstName string
		lastName string
		id int
	}
		yua.firstName = "tanemori"
		yua.lastName = "yua"
		yua.id = 779
	
	fmt.Println(yua)
}