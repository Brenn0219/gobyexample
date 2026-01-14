package examples

import (
	"fmt"
	"iter"
	"slices"
)

func (l *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := l.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}

			a, b = b, a+b
		}
	}
}

func RangeOverIterators() {
	l := List[int]{}
	l.Push(10)
	l.Push(13)
	l.Push(23)

	for e := range l.All() {
		fmt.Println(e)
	}

	all := slices.Collect(l.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
