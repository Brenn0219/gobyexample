package examples

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name, age: age}
	return &p
}

func Structs() {
	fmt.Println(person{"brenno", 21})
	fmt.Println(person{name: "alice", age: 12})
	fmt.Println(person{name: "bob"})
	fmt.Println(&person{name: "charlie", age: 30})
	fmt.Println(newPerson("david", 40))
	p := person{name: "eve", age: 25}
	fmt.Println(p.name, p.age)
	ps := &p
	fmt.Println(ps.age)

	dog := struct {
		name   string
		isGood bool
	}{
		name:   "pingo",
		isGood: true,
	}
	fmt.Println(dog)
}
