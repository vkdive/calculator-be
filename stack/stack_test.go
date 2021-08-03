package stack

import (
	"testing"

	"calculator/util"
	"github.com/stretchr/testify/assert"
)

func TestStack_ShouldReturnEmpty(t *testing.T) {
	stack := &Stack{Values: []util.Token{}}
	assert.Equal(t, stack.Length(), 0)
	assert.True(t, stack.IsEmpty())
}

func TestStack_PeekShouldReturnElementAtTop(t *testing.T) {
	stack := &Stack{Values: []util.Token{{
		Type:  util.OPERATOR,
		Value: "*",
	},{
		Type:  util.NUMBER,
		Value: "10",
	}}}
	assert.Equal(t, stack.Length(), 2)
	assert.False(t, stack.IsEmpty())
	peek := stack.Peek()
	assert.Equal(t, peek.Value, "10")
	assert.Equal(t, peek.Type, util.NUMBER)
}

func TestStack_PushShouldAddElementAtTop(t *testing.T) {
	stack := &Stack{Values: []util.Token{{
		Type:  util.OPERATOR,
		Value: "*",
	}}}
	stack.Push(util.Token{
		Type:  util.NUMBER,
		Value: "5",
	})
	assert.Equal(t, stack.Length(), 2)
	assert.False(t, stack.IsEmpty())
	peek := stack.Peek()
	assert.Equal(t, peek.Value, "5")
	assert.Equal(t, peek.Type, util.NUMBER)
}

func TestStack_PopShouldRemoveElementAtTop(t *testing.T) {
	stack := &Stack{Values: []util.Token{{
		Type:  util.OPERATOR,
		Value: "*",
	}}}
	stack.Push(util.Token{
		Type:  util.NUMBER,
		Value: "5",
	})
	assert.Equal(t, stack.Length(), 2)
	assert.False(t, stack.IsEmpty())
	peek := stack.Pop()
	assert.Equal(t, peek.Value, "5")
	assert.Equal(t, peek.Type, util.NUMBER)
}


