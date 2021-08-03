package scanner

import (
	"bufio"
	"bytes"
	"io"
	"unicode"

	"calculator/util"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return util.Eof
	}
	return ch
}

func (s *Scanner) Unread() {
	_ = s.r.UnreadRune()
}


func (s *Scanner) Scan() util.Token {
	ch := s.Read()

	if unicode.IsDigit(ch) {
		s.Unread()
		return s.ScanNumber()
	}  else if s.IsOperator(ch) {
		return util.Token{Type: util.OPERATOR, Value: string(ch)}
	} else if s.IsWhitespace(ch) {
		s.Unread()
		return s.ScanWhitespace()
	}

	switch ch {
	case util.Eof:
		return util.Token{Type: util.EOF}
	case '(':
		return util.Token{Type: util.LPAREN, Value: "("}
	case ')':
		return util.Token{Type: util.RPAREN, Value: ")"}
	}

	return util.Token{Type: util.ERROR, Value: string(ch)}
}

func (s *Scanner) ScanNumber() util.Token {
	var buf bytes.Buffer
	buf.WriteRune(s.Read())

	for {
		if ch := s.Read(); ch == util.Eof {
			break
		} else if !unicode.IsDigit(ch) && ch != '.' {
			s.Unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return util.Token{Type: util.NUMBER, Value: buf.String()}
}

func (s *Scanner) ScanWhitespace() util.Token {
	var buf bytes.Buffer
	buf.WriteRune(s.Read())

	for {
		if ch := s.Read(); ch == util.Eof {
			break
		} else if !s.IsWhitespace(ch) {
			s.Unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return util.Token{Type: util.WHITESPACE, Value: buf.String()}
}

func (s *Scanner)IsOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '^'
}

func (s *Scanner)IsWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

