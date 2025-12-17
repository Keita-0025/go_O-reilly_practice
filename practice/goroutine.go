package main 

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello")
}

func main () {
	// go sayHello()//①別スレッドで実行開始（またない）
	// fmt.Println("World")//②すぐにこっちを実行

	//一つのデータを送る
	// ch := make(chan string)

	// go func() {
	// 	ch <- "Hello"
	// }()
	// msg := <-ch
	// fmt.Println(msg)

	//複数のデータを送る
	// ch := make (chan int)
	// go func () {
	// 	ch <- 1
	// 	ch <- 3
	// 	ch <- 2
	// 	close(ch) 
	// }()

	// for num := range ch {
	// 	fmt.Println(num)
	// }

	ch := make(chan string) 

	go func () {
		defer close(ch)
		//「こんにちは」を一文字ずつおくる
		for _, char := range "こんにちは"{
			ch <- string(char)
		}
	}()

	//受け取って表示
	for text := range ch {
		fmt.Println(text)
	}

}