package main

import "fmt"

func prefixer(s string) func(string) string {
	greeting := s
	return func(name string)string {
		return greeting + " " + name
	} 
}


func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
