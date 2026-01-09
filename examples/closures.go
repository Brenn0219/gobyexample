package examples

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Closures() {
	nextInt := intSeq()

	println(nextInt())
	println(nextInt())
	println(nextInt())

	newInts := intSeq()
	println(newInts())
}
