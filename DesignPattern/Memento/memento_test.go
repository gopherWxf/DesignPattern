package Memento

import (
	"fmt"
	"testing"
)

func TestNumber_ReinstateMemento(t *testing.T) {
	n := NewNumber(7)
	n.Double()
	n.Double()
	memento := n.CreateMemento()
	n.Half()
	n.ReinstateMemento(memento)
	fmt.Println(n.value)
}
