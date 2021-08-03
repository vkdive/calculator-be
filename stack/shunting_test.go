package stack

import (
	"testing"

	"calculator/util"
	"github.com/stretchr/testify/assert"
)

func TestStack_ShuntingYardShouldReturnPostfixOperatorsForInfix(t *testing.T) {
	// expresiopn 3 *6
	// expected 36*
	stack := Stack{Values: []util.Token{{
		Type:  util.NUMBER,
		Value: "3",
	}, {
		Type:  util.OPERATOR,
		Value: "*",
	}, {
		Type:  util.NUMBER,
		Value: "6",
	}}}

	expected := Stack{Values: []util.Token{{
		Type:  util.NUMBER,
		Value: "3",
	}, {
		Type:  util.NUMBER,
		Value: "6",
	}, {
		Type:  util.OPERATOR,
		Value: "*",
	}}}
	assert.Equal(t, expected, stack.PerformShuntingYard())
}

func TestStack_ShuntingYardShouldReturnPostfixOperatorsForInfixWithBodmasRules(t *testing.T) {
	// expresiopn (3 *6)*5
	// expected 36*5*
	stack := Stack{Values: []util.Token{{
		Type:  util.LPAREN,
		Value: "(",
	}, {
		Type:  util.NUMBER,
		Value: "3",
	}, {
		Type:  util.OPERATOR,
		Value: "*",
	}, {
		Type:  util.NUMBER,
		Value: "6",
	}, {
		Type:  util.RPAREN,
		Value: ")",
	}, {
		Type:  util.OPERATOR,
		Value: "*",
	}, {
		Type:  util.NUMBER,
		Value: "5",
	}}}

	expected :=
		Stack{Values: []util.Token{{
			Type:  util.NUMBER,
			Value: "3",
		}, {
			Type:  util.NUMBER,
			Value: "6",
		}, {
			Type:  util.OPERATOR,
			Value: "*",
		}, {
			Type:  util.NUMBER,
			Value: "5",
		}, {
			Type:  util.OPERATOR,
			Value: "*",
		}}}

	assert.Equal(t, expected, stack.PerformShuntingYard())
}

func TestStack_ShuntingYardShouldReturnPostfixOperatorsForInfixWithOperatorPrecedence(t *testing.T) {
	// expresiopn 3*6+5
	// expected 36*5+
	stack := Stack{Values: []util.Token{ {
		Type:  util.NUMBER,
		Value: "3",
	}, {
		Type:  util.OPERATOR,
		Value: "*",
	}, {
		Type:  util.NUMBER,
		Value: "6",
	}, {
		Type:  util.OPERATOR,
		Value: "+",
	}, {
		Type:  util.NUMBER,
		Value: "5",
	}}}

	expected :=
		Stack{Values: []util.Token{{
			Type:  util.NUMBER,
			Value: "3",
		}, {
			Type:  util.NUMBER,
			Value: "6",
		}, {
			Type:  util.OPERATOR,
			Value: "*",
		}, {
			Type:  util.NUMBER,
			Value: "5",
		}, {
			Type:  util.OPERATOR,
			Value: "+",
		}}}

	assert.Equal(t, expected, stack.PerformShuntingYard())
}

func TestStack_SolvePostfixShouldReturnResultWithPostFixExpression(t *testing.T) {
	// expression 36*5+

	stack :=
		Stack{Values: []util.Token{{
			Type:  util.NUMBER,
			Value: "3",
		}, {
			Type:  util.NUMBER,
			Value: "6",
		}, {
			Type:  util.OPERATOR,
			Value: "*",
		}, {
			Type:  util.NUMBER,
			Value: "5",
		}, {
			Type:  util.OPERATOR,
			Value: "+",
		}}}

	assert.Equal(t, float64(23), stack.SolvePostfix())
}

