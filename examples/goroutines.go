package examples

import (
	"fmt"
	"time"
)

func g(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func Goroutines() {
	g("direct")

	go g("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
