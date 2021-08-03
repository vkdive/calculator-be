package parser

import (
	"strings"
	"testing"

	"calculator/stack"
	"calculator/util"
	"github.com/stretchr/testify/assert"
)

func TestParser_ParseShouldReturnStackInitializedWithExpression(t *testing.T) {
	parse, err := NewParser(strings.NewReader("3*6")).Parse()
	assert.NoError(t, err)
	expected := stack.Stack{Values: []util.Token{{
		Type:  util.NUMBER,
		Value: "3",
	}, {
		Type:  util.OPERATOR,
		Value: "*",
	}, {
		Type:  util.NUMBER,
		Value: "6",
	}}}
	assert.Equal(t, expected.Values, parse.Values)

}

func TestParser_ParseShouldReturnEmptyStackIfExpressionIsInvalidOrNotUnderstood(t *testing.T) {
	parse, _ := NewParser(strings.NewReader("MATH")).Parse()
	assert.Equal(t, 0, parse.Length())

}

