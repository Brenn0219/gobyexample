package examples

import "fmt"

type ServerState int

const (
	StateIDle ServerState = iota
	StateConected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIDle:     "idle",
	StateConected: "connected",
	StateError:    "error",
	StateRetrying: "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIDle:
		return StateConected
	case StateConected, StateRetrying:
		return StateIDle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func Enums() {
	ns := transition(StateIDle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}
