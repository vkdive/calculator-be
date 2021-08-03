package util

import "math"

func getOperators() map[string]struct {
	prec  int
	rAsoc bool
	fx    func(x float64, y float64) float64
} {
	return map[string]struct {
		prec  int
		rAsoc bool // true = right // false = left
		fx    func(x, y float64) float64
	}{
		"^": {4, true, func(x, y float64) float64 { return math.Pow(x, y) }},
		"*": {3, false, func(x, y float64) float64 { return x * y }},
		"/": {3, false, func(x, y float64) float64 { return x / y }},
		"+": {2, false, func(x, y float64) float64 { return x + y }},
		"-": {2, false, func(x, y float64) float64 { return x - y }},
	}
}


func GetPrecedence(val string, top string) bool{
	oprData:= getOperators()
	return (oprData[val].prec <= oprData[top].prec && oprData[val].rAsoc == false) ||
		(oprData[val].prec < oprData[top].prec && oprData[val].rAsoc == true)

}

func GetOperatorFunc(value string) func(x float64, y float64) float64 {
	oprData:= getOperators()
	return oprData[value].fx
}


