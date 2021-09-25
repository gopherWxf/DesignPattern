package Builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	builder:=NewConcreteBuilder()
	director:=NewDirector(&builder)
	director.Construct()
	result:=builder.GetResult()
	fmt.Println(result.built)
}
