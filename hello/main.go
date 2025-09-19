package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello,Go!")
	var w writer = ConsoleWriter{}
	w.write([]byte("hello world"))
}

type writer interface {
	write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
