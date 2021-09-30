package Interpreter

import "testing"
import "github.com/stretchr/testify/assert"

func TestInterpret(t *testing.T) {
	expression := "w x z +"
	sentence := NewEvaluator(expression)
	variable := make(map[string]Expression)
	variable["w"] = &Integer{6}
	variable["x"] = &Integer{10}
	variable["z"] = &Integer{41}
	result := sentence.Interpret(variable)
	assert.Equal(t, 51, result)

}
