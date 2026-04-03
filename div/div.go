package div

import "errors"

var errorDivisionByZero = errors.New("Division by zero")

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errorDivisionByZero
	}
	return a / b, nil
}
