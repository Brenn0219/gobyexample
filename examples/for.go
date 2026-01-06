package examples

import "fmt"

func For() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 4; j <= 6; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("rangei:", i)
	}

	for index, value := range []string{"a", "b", "c"} {
		fmt.Println("rangej:", index, value)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("odd:", n)
	}
}
