package main

import (
	"errors"
	"fmt"
)

var ErrDivisionByZero error = errors.New("division by zero")              // Деление на 0
var ErrMismatchedParentheses error = errors.New("mismatched parentheses") // Пропущена круглая скобка в выражении
var ErrInvalidExpression error = errors.New("invalid expression")         // Недопустимое выражение. Пропущен знак или цифра
var ErrEmptyExpression error = errors.New("empty expression")             // Пустое выражение

/*
Реализовать функцию func Calc(expression string) (float64, error) expression - строка-выражение состоящее из односимвольных идентификаторов и знаков арифметических действий Входящие данные - цифры(рациональные), операции +, -, *, /, операции приоритезации ( и ) В случае ошибки записи выражения функция выдает ошибку.

Сохраните этот код себе на github. Он понадобится вам при выполнении финальных заданий следующих модулей.
*/

// Функция возвращает результат заданной операции
//
// a, b float64 - переменные для вычисления
//
// operator rune - знак операции (+, -, *, /)
func calculate(a, b float64, operation rune) (float64, error) {
	switch operation {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil

	default:
		return 0, errors.New("this operation sign is unsupported. got: '%v', operation")
	}
}

// Проверка открывающих и закрывающих скобок
func checkParentheses(expression string) bool {
	return false
}

// Убрать пробелы
func simplifySpaces(inp string) string {
	var res string
	for _, v := range inp {
		if v != ' ' {
			res += string(v)
		}
	}
	return res
}

func Calc(expression string) (float64, error) {
	// Проверим что выражение не пустое
	expression = simplifySpaces(expression)
	exp := expression
	prio := [][]int
	l := len(expression)

	for i := 0; i < l; i++ {

		

	}






	// //Проверим корректность скобок
	// if checkParentheses(expression) == false {
	// 	return 0, ErrMismatchedParentheses
	// }

	// str := "(2+(2+2)*2)+2" // 12

	// Количество открытых скобок
	// parCount := 0
	// Слайс с индексами скобок --> [индекс открытия скобки]индекс закрытия скобки

	// var parOpen []int
	// var parClose []int
	// // lastInd := 0
	// for i := 0; i < len(expression); i++ {
	// 	cur := expression[i]

		// if cur == '(' || cur == ')' {

		// 	if cur == ')' && len(parOpen) != 0 /*parCount == 0 */ {
		// 		return 0, ErrMismatchedParentheses

		// 	} else if cur == '(' {
		// 		// parCount++
		// 		parOpen = append(parOpen, i)
		// 		continue

		// 	} else if cur == ')' {
		// 		parClose = append(parClose, i)
		// 	}
		// }
	}

	return 0, nil
}

func main() {
	str := "1+(5+(2+1)*3)+7" // 22
	/*

	

	*/
	fmt.Println(Calc(str))
}
