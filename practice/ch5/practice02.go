package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	data := make([]byte, 2048)
	total := 0
	for {
		count, err := f.Read(data)
		total += count
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, nil
		}
	}
	return total, nil
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	count, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(count)
}
