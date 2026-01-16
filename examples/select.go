package examples

import (
	"fmt"
	"time"
)

func Select() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
