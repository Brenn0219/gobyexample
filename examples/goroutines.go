package examples

import (
	"fmt"
	"time"
)

func h(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func Goroutines() {
	h("direct")

	go h("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
