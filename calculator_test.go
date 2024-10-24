package main

import (
	"errors"
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
		{"3+r-14+ghrkfjd01", 0},
		{"234/0*2423", 0},
		{"534%432&23!21", 0},
		{"32++23443-1212*2", 0},
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
		id         string
		expression string
		want       float64
		err        error
	}{
		{"1", "2+2", 4, nil},
		{"2", "2*2", 4, nil},
		{"3", "2/2", 1, nil},
		{"4", "2-2", 0, nil},
		{"5", "2+2-2", 2, nil},
		{"6", "2+2*2", 6, nil},
		{"7", "2+2/2", 3, nil},
		{"8", "2*3+3*2-5*5", -13, nil},
		{"9", "(3+5)*2", 16, nil},
		{"10", "10/(2+3)", 2, nil},
		{"11", "3.5*2+1", 8, nil},
		{"12", "1+2*(3+4)", 15, nil},
		{"13", "(1+(4+5+2)-3)+(6+8)", 23, nil},
		{"14", "15/0", 0, ErrDivisionByZero},
		{"15", "(2+3", 0, ErrMismatchedParentheses},
		{"16", "2++2", 0, ErrInvalidExpression},
		{"17", "2*2)", 0, ErrMismatchedParentheses},
		{"18", "1+2*3-", 0, ErrInvalidExpression},
		{"19", "", 0, ErrEmptyExpression},
		{"20", "(8/4+5-2)/10", 0.5, nil},
		{"21", "(((2*2))", 0, ErrMismatchedParentheses},
		{"22", "()(2*2)", 0, ErrInvalidExpression},
		{"23", "(((2*2))", 0, ErrMismatchedParentheses},
		{"24", ")()2*2))", 0, ErrMismatchedParentheses},
		{"25", "((2*2)))", 0, ErrMismatchedParentheses},
		{"26", "2*2(", 0, ErrMismatchedParentheses},
		{"27", ")2*2)", 0, ErrMismatchedParentheses},
		{"28", ")2*2", 0, ErrMismatchedParentheses},
		{"29", "     5+4  -4-2", 3, nil},
		{"30", "2+2+2+2", 8, nil},
		{"31", "2+2*2", 6, nil},
		{"32", "2*2-2*2", 0, nil},
		{"33", "2*2/4*2*51", 102, nil},
		{"34", "2/4", 0.5, nil},
		{"35", "2+2-5+6*34/10-20+5.59", 4.99, nil},
		{"36", "3+r-14+ghrkfjd01", 0, ErrInvalidExpression},
		{"37", "234/0*2423", 0, ErrDivisionByZero},
		{"38", "534%432&23!21", 0, ErrInvalidExpression},
		{"39", "32++23443-1212*2", 0, ErrInvalidExpression},
		{"40", "2", 2, nil},
		{"41", "+", 0, ErrInvalidExpression},
		{"42", "*", 0, ErrInvalidExpression},
		{"43", "1+(5+(2+1)*3)+7", 22, nil},
		{"44", "(0)*(2*2))", 0, ErrMismatchedParentheses},
		{"45", "(2)+(3)+1", 6, nil},
		{"46", "10*10-1", 99, nil},
		{"47", "255*34-166*2", 8338, nil},
	}

	for _, tc := range cases {
		got, err := Calc(tc.expression)
		if !errors.Is(err, tc.err) {
			t.Errorf("ERROR: id:%v -> Calc(%v) -> wrong err! want err: %v; got err: %v; want val: %v; got val: %v", tc.id, tc.expression, tc.err, err, tc.want, got)
		} else if err != nil && !errors.Is(err, tc.err) {
			t.Errorf("FATAL!!!: id:%v -> Calc(%v) -> unexpected error!! want err: %v; got err: %v; want val: %v; got val: %v", tc.id, tc.expression, tc.err, err, tc.want, got)
		}

		temp, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", got), 64)

		if temp != tc.want {
			t.Errorf("WA: id:%v -> Calc(%v): want: %v; got: %v", tc.id, tc.expression, tc.want, got)
		}
	}
}
