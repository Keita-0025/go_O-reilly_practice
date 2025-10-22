package main

import "fmt"

func main() {
	/**
	** map[string]intはstringをkeyにintをvalueに持つmap
	**/
	// m := map[string]int{
	// 	"apple": 1,
	// 	"banana": 2,
	// 	"cherry": 3,
	// }
	// v, ok := m["apple"]
	// fmt.Println(v, ok) // 1 true

	// v, ok = m["pineapple"]
	// fmt.Println(v, ok) // 0 false

	// v, ok = m["banana"]
	// fmt.Println(v, ok) // 2 true

	// v, ok = m["cherry"]
	// fmt.Println(v, ok) // 3 true

	/**
	** map[int]boolはintをkeyにboolをvalueに持つmap
	**/
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10,100}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet)) // 12 8
	fmt.Println(intSet[5]) // true
	fmt.Println(intSet[500]) // false
	if intSet[100] {
		fmt.Println("100は含まれている")
	}

	/**
	** map[int]struct{}はintをkeyに空の構造体をvalueに持つmap
	**/
	intSet2 := map[int]struct{}{}
	vals2 := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10,100}
	for _, v := range vals2 {
		intSet2[v] = struct{}{}
	}
	fmt.Println(len(vals), len(intSet2)) // 12 8
	fmt.Println(intSet2[5]) // {}
	fmt.Println(intSet2[500]) // {}
	if _, ok := intSet2[5]; ok {
		fmt.Println("5は含まれている")
	}
	if _, ok := intSet2[500]; ok {
		fmt.Println("500は含まれている")
	}
}
