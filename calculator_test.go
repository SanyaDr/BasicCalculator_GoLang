package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCalculateSimple(t *testing.T) {
	cases := []struct {
		inp  string
		want float64
	}{
		{"2+2+2+2", 8},
		{"2+2*2", 6},
		{"2*2-2*2", 0},
		{"2*2/4*2*51", 102},
		{"2/4", 0.5},
		{"2+2-5+6*34/10-20+5.59", 4.99},
	}

	for _, c := range cases {
		got, _ := calculateSimple(c.inp)
		temp, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", got), 64)
		if temp != c.want {
			t.Errorf("simplifyParentheses(%v);  -->   got: %v; want %v", c.inp, got, c.want)
		}
	}
}

func TestParentheses(t *testing.T) {
	cases := []struct {
		inp  string
		want string
	}{
		{"1+(5+(2+1)*3)+7", "1+14+7"},
	}

	for _, c := range cases {
		got, _ := simplifyParentheses(c.inp)
		if got != c.want {
			t.Errorf("simplifyParentheses(%q);  -->   got: %q; want %q", c.inp, got, c.want)
		}
	}
}

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
