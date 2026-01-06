package examples

import "time"

func Switch() {
	i := 2
	switch i {
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	default:
		println("other")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend")
	default:
		println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("It's before noon")
	default:
		println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			println("I'm a bool")
		case int:
			println("I'm an int")
		default:
			println("Don't know type")
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
