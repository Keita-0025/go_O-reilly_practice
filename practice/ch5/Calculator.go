package main

import (
	"fmt"
	"strconv"
)

func add (i,j int)int{return i+j}
func sub (i,j int)int{return i-j}
func mul (i,j int)int{return i*j}
func div (i,j int)int{return i/j}

func main() {

	type opFuncType func(int,int)int
	
	var opMap = map[string]opFuncType {
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
	
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"tow", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
	}
	
	for _, expression := range expressions {
		if len(expression) != 3 { //演算子と被演算子の合計個数をチェック
			fmt.Print(expression,"--不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok{
			fmt.Println(expression, " -- ", "定義されていない演算子です:", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		result := opFunc(p1, p2)//実際の計算
		fmt.Print(expression, " → ", result, "\n")
	}
	main2()
	main3()
}

func main2 () {
	f := func (j int) {
		fmt.Println("無名関数の中で", j, "を出力")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}
}

func main3() {
	for i := 5; i < 10; i++ {
		func (j int) {
			fmt.Println("無名関数の中で", j, "を出力")
		}(i)
	}
}
