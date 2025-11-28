package main 

import (
	"fmt"
)

type MyFuncOpts struct {
	FirstName string
	LastName string
	Age int
}
func MyFunc(opts MyFuncOpts) error {
	fmt.Println(opts)
	fmt.Println("<ここで必要な処理を行う>")
	return nil
}

// func main () {
// 	MyFunc(MyFuncOpts{
// 		LastName: "Patel",
// 		Age: 50,
// 	})
// 	MyFunc(MyFuncOpts{
// 		FirstName: "Joe",
// 		LastName: "Smith",
// 	})
// }

func addTo (base int, vals ...int)[]int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// func main() {
// 	fmt.Println(addTo(3))
// 	fmt.Println(addTo(3,2))
// 	fmt.Println(addTo(3,2,1,4))
// 	a:=[]int{2,3}
// 	fmt.Println(addTo(3, a...))
// 	fmt.Println(addTo(3, []int{1,2,3,4,5}...))
// 	// 可変長の引数を渡す時は...をつけないとコンパイラエラーになる
// 	// fmt.Println(addTo(3,a))
// 	// fmt.Println(addTo(3, []int{1,2,3,4,5}))
// }
