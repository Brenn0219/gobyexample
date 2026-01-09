package examples

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func Functions() {
	res := plus(1, 2)
	println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	println("1 + 2 + 3 =", res)
}
