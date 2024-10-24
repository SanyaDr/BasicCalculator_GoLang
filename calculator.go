package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrDivisionByZero error = errors.New("division by zero")                         // Деление на 0
var ErrMismatchedParentheses error = errors.New("mismatched parentheses")            // Пропущена круглая скобка в выражении
var ErrInvalidExpression error = errors.New("invalid expression")                    // Недопустимое выражение. Пропущен знак или цифра
var ErrEmptyExpression error = errors.New("empty expression")                        // Пустое выражение
var ErrUnsupportedOperation error = errors.New("this operation sign is unsupported") // Данная операция не поддерживается

var operatorChars = "+-*/" // Поддерживаемые операции
var separatorOperator = func(r rune) bool {
	return strings.ContainsRune(operatorChars, r)
}

// Возвращает слайс чисел из строки
func separateNumbers(inp string) ([]float64, error) {
	numbersStr := strings.FieldsFunc(inp, separatorOperator)
	numbers := make([]float64, len(numbersStr))
	for i, s := range numbersStr {
		var err error
		numbers[i], err = strconv.ParseFloat(s, 64)
		if err != nil {
			return numbers, ErrInvalidExpression
		}
	}
	return numbers, nil
}

// Возвращает слайс операторов из строки
func separateOperators(inp string) []rune {
	var operators []rune
	for _, v := range inp {
		if strings.ContainsRune(operatorChars, v) {
			operators = append(operators, v)
		}
	}
	return operators
}

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
		return 0, ErrUnsupportedOperation
	}
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

// Вычисляет хуйню без скобок, пример: 2+3*4/5
func calculateSimple(inp string) (float64, error) {
	numbers, err := separateNumbers(inp) // Слайс чисел
	operators := separateOperators(inp)  // Слайс операторов
	if err != nil {
		return 0, err
	}
	// Если операторов больше чем чисел, то ошибка
	if len(operators) >= len(numbers) {
		return 0, ErrInvalidExpression
	}

	// Если есть приоритетные операторы (умножения, деления)
	if strings.ContainsAny(inp, "*/") {
		for i := 0; i < len(operators); i++ {
			if operators[i] == '*' || operators[i] == '/' {
				numbers[i], err = calculate(numbers[i], numbers[i+1], operators[i])
				if err != nil {
					return 0, err
				}
				numbers = append(numbers[:i+1], numbers[i+2:]...)
				operators = append(operators[:i], operators[i+1:]...)
				i--
			}
		}
	}
	//если приоритетных операций больше нет
	for len(numbers) != 1 {
		numbers[0], err = calculate(numbers[0], numbers[1], operators[0])
		if err != nil {
			return 0, err
		}
		numbers = append(numbers[:1], numbers[2:]...)
		operators = operators[1:]
	}
	return numbers[0], nil
}

func simplifyParentheses(expression string) (string, error) {
	parOpenCount, parCloseCount := 0, 0
	parFirstOpen, parLastClose := 0, 0
	// ищем срез внутри скобок ->
	// если скобок нет то кидаем хуйню в s1mple
	// возвращаем ответ

	for i := 0; i < len(expression); i++ {
		cur := expression[i]
		if cur == '(' {
			parOpenCount++
			if parOpenCount == 1 {
				parFirstOpen = i
			}
		}
		if cur == ')' {
			parCloseCount++
			if parCloseCount == parOpenCount {
				parLastClose = i

				simpleExp, err := simplifyParentheses(expression[parFirstOpen+1 : parLastClose])
				if err != nil {
					return "", ErrInvalidExpression
				}
				num, err := calculateSimple(simpleExp)
				if err != nil {
					return "", err
				}
				//i -= len(expression) - len(simpleExp)
				i -= parLastClose - parFirstOpen
				expression = expression[:parFirstOpen] + fmt.Sprintf("%v", num) + expression[parLastClose+1:]

				parOpenCount, parCloseCount = 0, 0
			}
		}
	}
	if parOpenCount != parCloseCount {
		return "", ErrMismatchedParentheses
	}
	//if parOpenCount > 0 {
	//	simpleExp, err := simplifyParentheses(expression[parFirstOpen+1 : parLastClose])
	//	if err != nil {
	//		return "", ErrInvalidExpression
	//	}
	//	num, err := calculateSimple(simpleExp)
	//	if err != nil {
	//		return "", err
	//	}
	//
	//	expression = expression[:parFirstOpen] + fmt.Sprintf("%v", num) + expression[parLastClose+1:]
	//}

	return expression, nil
}

func Calc(expression string) (float64, error) {
	// Проверим что выражение не пустое
	if len(expression) == 0 {
		return 0, ErrEmptyExpression
	}
	expression = simplifySpaces(expression)

	simple, err := simplifyParentheses(expression)
	if err != nil {
		return 0, err
	}
	res, err := calculateSimple(simple)
	if err != nil {
		return 0, err
	}

	return res, nil
}
