package scanner

import (
	"strings"
	"testing"

	"calculator/util"
	"github.com/stretchr/testify/assert"
)


func TestScanner_IsOperatorShouldReturnTrueOrFalse(t *testing.T) {
	expected := NewScanner(nil).IsOperator('*')
	assert.True(t, expected)
}

func TestScanner_IsWhiteSpaceShouldReturnTrueOrFalse(t *testing.T) {
	expected := NewScanner(nil).IsWhitespace(' ')
	assert.True(t, expected)
}

func TestScanner_ScanShouldReturnNumberTokenForGivenRune(t *testing.T) {
	token := NewScanner(strings.NewReader("3")).Scan()
	assert.Equal(t, util.Token{
		Type:  util.NUMBER,
		Value: "3",
	}, token)
}

func TestScanner_ScanShouldReturnParenthesisTokenForGivenRune(t *testing.T) {
	token := NewScanner(strings.NewReader("(")).Scan()
	assert.Equal(t, util.Token{
		Type:  util.LPAREN,
		Value: "(",
	}, token)
}