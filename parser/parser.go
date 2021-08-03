package parser

import (
	"fmt"
	"io"

	"calculator/scanner"
	"calculator/stack"
	"calculator/util"
)

type Parser struct {
	s   *scanner.Scanner
	buf struct {
		tok util.Token
		n   int
	}
	stack *stack.Stack
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: scanner.NewScanner(r)}
}

func (p *Parser) Scan() (tok util.Token) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok
	}

	tok = p.s.Scan()

	p.buf.tok = tok

	return
}

func (p *Parser) ScanIgnoreWhitespace() (tok util.Token) {
	tok = p.Scan()
	if tok.Type == util.WHITESPACE {
		tok = p.Scan()
	}
	return
}

func (p *Parser) Unscan() {
	p.buf.n = 1
}

func (p *Parser) Parse() (stack.Stack, error) {
	stack := stack.New()
	for {
		tok := p.ScanIgnoreWhitespace()
		if tok.Type == util.ERROR {
			return stack, fmt.Errorf("ERROR: %q", tok.Value)
		} else if tok.Type == util.EOF {
			break
		} else if tok.Type == util.OPERATOR && tok.Value == "-" {
			lastTok := stack.Peek()
			nextTok := p.ScanIgnoreWhitespace()
			if (lastTok.Type == util.OPERATOR || lastTok.Value == "" || lastTok.Type == util.LPAREN) && nextTok.Type == util.NUMBER {
				stack.Push(util.Token{Type: util.NUMBER, Value: "-" + nextTok.Value})
			} else {
				stack.Push(tok)
				p.Unscan()
			}
		} else {
			stack.Push(tok)
		}
	}
	return stack, nil
}
