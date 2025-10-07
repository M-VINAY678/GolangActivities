package main

import (
	"fmt"
	"slices"
)

type NumberOrString interface {
	int | float64 | string
}
type Queue[T NumberOrString] []T

func (q *Queue[T]) enqueue(value T) {
	*q = append(*q, value)
}
func (q *Queue[T]) dequeue() (err error) {
	if len((*q)) == 0 {
		return fmt.Errorf("Queue is Empty, can't dequeue it")
	}
	*q = slices.Delete(*q, 0, 1)
	return err
}
func (q *Queue[T]) peek() (res T, err error) {
	if len((*q)) == 0 {
		return res, fmt.Errorf("Queue is Empty,can't use peek method")
	}
	return (*q)[0], err
}

func main() {
	var intQueue Queue[int]
	intQueue.enqueue(5)
	intQueue.enqueue(6)
	intQueue.dequeue()
	intQueue.enqueue(4)
	intQueue.dequeue()
	intQueue.dequeue()
	err := intQueue.dequeue()
	if err != nil {
		fmt.Println(err)
	}
	peek, err := intQueue.peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The Peek of intQueue is : ", peek)
	fmt.Println("intQueue : ", intQueue)
	var stringQueue Queue[string]
	stringQueue.enqueue("Hello")
	stringQueue.enqueue("sample")
	stringQueue.dequeue()
	stringQueue.enqueue("world")
	stringPeek, _ := stringQueue.peek()
	fmt.Println("The Peek of stringQueue is : ", stringPeek)
	fmt.Println("stringQueue : ", stringQueue)

}
