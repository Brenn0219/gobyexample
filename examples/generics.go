package examples

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}

	return -1
}

type node[T any] struct {
	next *node[T]
	val  T
}

type List[T any] struct {
	head, tail *node[T]
}

func (l *List[T]) Push(v T) {
	if l.tail == nil {
		l.head = &node[T]{val: v}
		l.tail = l.head
	} else {
		l.tail.next = &node[T]{val: v}
		l.tail = l.tail.next
	}
}

func (l *List[T]) AllElements() []T {
	var elems []T

	for e := l.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}

	return elems
}

func Generics() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo: ", SlicesIndex(s, "zoo"))

	l := List[int]{}
	l.Push(10)
	l.Push(11)
	l.Push(22)
	fmt.Println("list: ", l.AllElements())
}
