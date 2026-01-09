package examples

func vals() (int, int) {
	return 3, 7
}

func MultipleReturnValues() {
	a, b := vals()
	println("a:", a, "b:", b)
	_, c := vals()
	println("c:", c)
}
