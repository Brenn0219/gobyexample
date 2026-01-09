package examples

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func Recursion() {
	println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	println(fib(7))
}
