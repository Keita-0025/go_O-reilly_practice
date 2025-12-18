package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add2(i, j int) (int, error) { return i + j, nil }
func sub2(i, j int) (int, error) { return i - j, nil }
func mul2(i, j int) (int, error) { return i * j, nil }
func div2(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("0で割ることはできません")
	}
	return i / j, nil
}

type opFuncType func(int, int) (int, error)

func main() {
	var opMap = map[string]opFuncType{
		"+": add2,
		"-": sub2,
		"*": mul2,
		"/": div2,
	}

	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"2", "/", "0"},
		[]string{"two", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Print(expression, "--不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println(expression, " -- ", "定義されていない演算子です:", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		fmt.Print(expression, " → ", result, "\n")
	}
}


