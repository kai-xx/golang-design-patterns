package Prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 原型模式
func TestConcretePrototype_Clone(t *testing.T) {
	name := "出去浪"
	proto := ConcretePrototype{
		name: name,
	}
	newProto := proto.Clone()
	actualName := newProto.Name()

	assert.Equal(t, name, actualName)
}
