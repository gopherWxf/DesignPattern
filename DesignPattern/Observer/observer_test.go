package Observer

import (
	"sync"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	n := ConcreteSubject{Observers: sync.Map{}}
	obs1 := ConcreteObserver{
		ID:   1,
		Time: time.Now(),
	}
	obs2 := ConcreteObserver{
		ID:   1,
		Time: time.Now(),
	}
	n.AddListener(&obs1)
	n.AddListener(&obs2)
	for x := range Fib(10) {
		n.Notify(Event{Data: x})
	}
}
