package util


type Token struct {
	Type  int
	Value string
}

var Eof = rune(0)

const (
	NUMBER int = iota
	LPAREN
	RPAREN
	OPERATOR
	WHITESPACE
	ERROR
	EOF
)