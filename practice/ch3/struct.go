package main

import "fmt"

/**
** type person struct(構造体)はpersonという型を定義する
**/
func main() {

	/**
	** personはname, age, petをファイルに持つ
	**/
	type person struct {
		name string
		age int
		pet string
	}

	keita := person{
		"けいた",
		29,
		"キャット",
	}
	fmt.Println(keita) // {けいた 29 キャット}

	mayu := person{
		age: 29,
		name: "まゆ",
	}
	fmt.Println(mayu) // {まゆ 29 }

	keita.name = "keita" 
	fmt.Println(keita.name) // {keita 29 キャット}

	var user struct {
		name string
		age int
		from string
		role string
	}

	user.name = "きよと"
	user.age = 24
	user.from = "japan"
	user.role = "developer"

	role := struct {
		name string
		level int
	}{
		name: "backend",
		level: 10,
	}
	fmt.Println(role) // {backend 10}


	type firstPerson struct {
		name string
		age int
	}

	type secondPerson struct {
		name string
		age int
	}
	
	type thirdPerson struct {
		age int
		name string
	}
	type fourthPerson struct {
		firstName string
		age int
	}
	type fifthPerson struct {
		name string
		age int
		favoriteColor string
	}

	f := firstPerson{
		name: "太郎",
		age: 24,
	}
	var g struct {
		name string
		age int
	}

	g = f
	fmt.Println(g)
	fmt.Println(g == f) // {太郎 24}


}
