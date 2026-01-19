package examples

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker2(i)
		})
	}

	wg.Wait()
}
