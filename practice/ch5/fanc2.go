package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		return result, remainder, errors.New("0で割ることはできません")
	}
	result, remainder = num/denom, num%denom
	return result, remainder, nil
	//ブランクリターンは戻り値が分からなくなるから戻り値を返す関数には絶対に使わない
}

func main() {
	rs, rm, _ := divAndRemainder(5, 2)
	fmt.Println(rs, rm)
	main2()
}

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	//for-rangeでループするとvはrune(int32)型になる　ASCII値の合計
	for _, v := range a {
		total += int(v)
	}
	return total
}

func main2() {
	var myFuncVariable func(string) int
	myFuncVariable = f1
	result := myFuncVariable("Hello")
	fmt.Println(result) //5

	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println(result)
}
