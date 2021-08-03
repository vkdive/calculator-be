package IntegretionTests

import (
	"strings"
	"testing"

	"calculator/parser"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnValueAfterEvaluating(t *testing.T) {
	p := parser.NewParser(strings.NewReader("(3*5)*6/4(2*6)"))
	stack, _ := p.Parse()
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, float64(90), stack.PerformShuntingYard().SolvePostfix())
}