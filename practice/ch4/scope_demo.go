package main

import "fmt"

func main() {
	// デモ1: シャドウイングの動作
	fmt.Println("=== デモ1: シャドウイング ===")
	var total int
	fmt.Printf("ループ前: total = %d (アドレス: %p)\n", total, &total)

	for i := 0; i < 3; i++ {
		// ここで新しいtotalを宣言（シャドウイング）
		total := i + total // 外側のtotalの値（0）を使用
		fmt.Printf("ループ内 i=%d: total = %d (アドレス: %p)\n", i, total, &total)
	}

	fmt.Printf("ループ後: total = %d (アドレス: %p)\n", total, &total)
	fmt.Println()

	// デモ2: 正しい実装（代入を使用）
	fmt.Println("=== デモ2: 正しい実装（代入） ===")
	var total2 int
	fmt.Printf("ループ前: total2 = %d\n", total2)

	for i := 0; i < 3; i++ {
		total2 = i + total2 // := ではなく = を使用（代入）
		fmt.Printf("ループ内 i=%d: total2 = %d\n", i, total2)
	}

	fmt.Printf("ループ後: total2 = %d\n", total2)
}

