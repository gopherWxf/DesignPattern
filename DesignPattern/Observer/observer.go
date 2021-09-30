package Observer

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Data int
}

type Observer interface {
	NotifyCallBack(event Event)
}

type Subject interface {
	AddListener(observer Observer)
	RemoveListener(observer Observer)
	Notify(event Event)
}

type ConcreteObserver struct {
	ID   int
	Time time.Time
}

type ConcreteSubject struct {
	Observers sync.Map
}

func (e *ConcreteObserver) NotifyCallBack(event Event) {
	fmt.Println(fmt.Sprintf("Recieved:%d after %v\n", event.Data, time.Since(e.Time)))
}

func (e *ConcreteSubject) AddListener(obs Observer) {
	e.Observers.Store(obs, struct{}{})
}

func (e *ConcreteSubject) RemoveListener(obs Observer) {
	e.Observers.Delete(obs)
}

func (e *ConcreteSubject) Notify(event Event) {
	e.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallBack(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = j, i+j {
			out <- i
		}
	}()
	return out
}
