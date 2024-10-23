package main

import (
	"testing"
)

func TestCalc(t *testing.T) {

	cases := []struct {
		expression string
		want       float64
		err        error
	}{
		{"2+2", 4, nil},
		{"2*2", 4, nil},
		{"2/2", 1, nil},
		{"2-2", 0, nil},
		{"2+2-2", 2, nil},
		{"2+2*2", 6, nil},
		{"2+2/2", 3, nil},
		{"2*3+3*2-5*5", 37, nil},
		{"(3+5)*2", 16, nil},
		{"10/(2+3)", 2, nil},
		{"3.5*2+1", 8, nil},
		{"1+2*(3+4)", 15, nil},
		{"(1+(4+5+2)-3)+(6+8)", 23, nil},
		{"15/0", 0, ErrDivisionByZero},
		{"(2+3", 0, ErrMismatchedParentheses},
		{"2++2", 0, ErrInvalidExpression},
		{"2*2)", 0, ErrMismatchedParentheses},
		{"1+2*3-", 0, ErrInvalidExpression},
		{"", 0, ErrEmptyExpression},
		{"(8/4+5-2)/10", 0.5, nil},
		{"(((2*2))", 0, ErrMismatchedParentheses},
		{"()(2*2))", 0, ErrMismatchedParentheses},
		{"(((2*2))", 0, ErrMismatchedParentheses},
		{")()2*2))", 0, ErrMismatchedParentheses},
		{"((2*2)))", 0, ErrMismatchedParentheses},
		{"2*2(", 0, ErrMismatchedParentheses},
		{")2*2)", 0, ErrMismatchedParentheses},
		{")2*2", 0, ErrMismatchedParentheses},
		{"     5+4  -4-2", 3, nil},
	}

	for _, tc := range cases {
		got, err := Calc(tc.expression)
		if err != tc.err {
			t.Errorf("ERROR: Calc(%v) -> wrong err! want err: %v; got err: %v; want val: %v; got val: %v", tc.expression, tc.err, err, tc.want, got)
		}
		if got != tc.want {
			t.Errorf("ERROR: Calc(%v): want: %v; got: %v", tc.expression, tc.want, got)
		}
	}
}
