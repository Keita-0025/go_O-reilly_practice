package main

import "fmt"

func main() {
	/**
	** x,y,zは同じ配列を参照しているから値を変えると全て変わる
	**/
	// x := []string{"a", "b", "c", "d"}
	// y := x[:2]
	// z := x[1:]
	// x[1] = "y"
	// y[0] = "x"
	// z[1] = "z"
	// fmt.Println(x) // [a y c d]
	// fmt.Println(y) // [a y]
	// fmt.Println(z) // [y c d]

	/**
	** yにappendするとxのキャパシティが増える
	**/
	// x := [] string{"a", "b", "c", "d"}
	// y := x[:2] // 先頭から2つ
	// y = append(y, "z")
	// y = append(y, "1")
	// y = append(y, "2")
	// fmt.Println(x) // [a b z 2]
	// fmt.Println(y) // [a b z]

	/**
	** yにappendするとxのキャパシティが増える
	**/
	// x := make([]string, 0, 5)
	// x = append(x, "a", "b", "c", "d")
	// y := x[:2] // 先頭から2つ
	// z := x[2:] // 2つめから残り全部
	// fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3
	// y = append(y, "i", "j", "k")
	// fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3
	// y = append(y, "x")
	// fmt.Println(cap(x), cap(y), cap(z)) // 5 10 3
	// y = append(y, "y")
	// fmt.Println(cap(x), cap(y), cap(z)) // 5 10 3
	// fmt.Println(x) // [a b i j]
	// fmt.Println(y) // [a b i j k x y]
	// fmt.Println(z) // [i j]

	/**
	** yにappendしてもxのキャパシティは増えない
	**/
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2] //yが利用できるのはxの2番目まで
	z := x[2:4:4] //zが利用できるのはxの4番目まで
	fmt.Println(cap(x), cap(y), cap(z)) // 5 2 2
	y = append(y, "i", "j", "k")
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 2
	y = append(y, "x")
	fmt.Println(cap(x), cap(y), cap(z)) // 5 10 2
	y = append(y, "y")
	fmt.Println(cap(x), cap(y), cap(z)) // 5 10 2
	fmt.Println(x) // [a b c d]
	fmt.Println(y) // [a b i j k x y]
	fmt.Println(z) // [c d]
}