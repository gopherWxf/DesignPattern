package Prototype

import "testing"
import "github.com/stretchr/testify/assert"
func TestConcretePrototype_Clone(t *testing.T) {
	name:="wxf"
	p:=ConcretePrototype{name: name}
	newProto:=p.Clone()
	assert.Equal(t,name,newProto.Name())
}