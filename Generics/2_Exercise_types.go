package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Append 為什麼 List[T] 前面要加 * 號？
// 為什麼是 List[T] 而不是 List ?
// newNode := &List[T]{val: value} 這行為什麼要加 & ?
func (l *List[T]) Append(value T) *List[T] {
	newNode := &List[T]{val: value}
	if l == nil {
		return newNode
	}

	// Find the last node
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return l
}

func (l *List[T]) PrintAll() {
	current := l
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func main() {
	var list *List[int] // Create a list for integers
	list = list.Append(10)
	list = list.Append(20)
	list = list.Append(30)
	list.PrintAll() // This should print 10, 20, 30 each on a new line

	var strList *List[string] // Create a list for strings
	strList = strList.Append("hello")
	strList = strList.Append("world")
	strList.PrintAll() // This should print "hello" and "world"
}
