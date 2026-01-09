package examples

import "fmt"

func RangeOverBuiltInTypes() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			println("index of 3 is", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
