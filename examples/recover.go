package examples

import "fmt"

func myPanic() {
	panic("a problem")
}

func Recover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered. Error", r)
		}
	}()

	myPanic()

	fmt.Println("After myPanic")
}
