package stack

import (
	"strconv"

	"calculator/util"
)

func (s Stack) PerformShuntingYard() Stack {
	postfix := New()
	operators := New()
	for _, v := range s.Values {
		switch v.Type {
		case util.OPERATOR:
			for !operators.IsEmpty() {
				val := v.Value
				top := operators.Peek().Value
				if util.GetPrecedence(val, top) {
					postfix.Push(operators.Pop())
					continue
				}
				break
			}
			operators.Push(v)
		case util.LPAREN:
			operators.Push(v)
		case util.RPAREN:
			for i := operators.Length() - 1; i >= 0; i-- {
				if operators.Values[i].Type != util.LPAREN {
					postfix.Push(operators.Pop())
					continue
				} else {
					operators.Pop()
					break
				}
			}
		default:
			postfix.Push(v)
		}
	}
	operators.EmptyInto(&postfix)
	return postfix
}

func (s Stack) SolvePostfix() float64 {
	stack := New()
	for _, v := range s.Values {
		switch v.Type {
		case util.NUMBER:
			stack.Push(v)
		case util.OPERATOR:
			f := util.GetOperatorFunc(v.Value)
			var x, y float64
			y, _ = strconv.ParseFloat(stack.Pop().Value, 64)
			x, _ = strconv.ParseFloat(stack.Pop().Value, 64)
			result := f(x, y)
			stack.Push(util.Token{Type: util.NUMBER, Value: strconv.FormatFloat(result, 'f', -1, 64)})
		}
	}
	out, _ := strconv.ParseFloat(stack.Values[0].Value, 64)
	return out
}
