package Factory_Method

import (
	"fmt"
	"testing"
)

func TestCreateOperation(t *testing.T) {
	fac := &AddFactory{}
	opAdd := fac.CreateOperation()
	opAdd.SetA(1)
	opAdd.SetB(2)
	fmt.Println(opAdd.GetResult())
}
