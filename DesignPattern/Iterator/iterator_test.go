package Iterator

import (
	"fmt"
	"testing"
)

func TestArrayIterator_Value(t *testing.T) {
	array := []interface{}{6,8,7,2,5,0,3,2}
	a := 0
	iterator := ArrayIterator{array: array, index: &a}
	for it := iterator; iterator.HasNext(); iterator.Next() {
		index, value := it.Index(), it.Value().(int)
		for value != array[*index].(int) {
			fmt.Println("error")
		}
		fmt.Println(*index, value)
	}
}
