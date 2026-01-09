package examples

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Pointers() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("after zeroval:", i)

	zeroptr(&i)
	fmt.Println("after zeroptr:", i)
}
