package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrDivisionByZero error = errors.New("division by zero")              // Деление на 0
var ErrMismatchedParentheses error = errors.New("mismatched parentheses") // Пропущена круглая скобка в выражении
var ErrInvalidExpression error = errors.New("invalid expression")         // Недопустимое выражение. Пропущен знак или цифра
var ErrEmptyExpression error = errors.New("empty expression")             // Пустое выражение

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
			return numbers, err
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
		return 0, errors.New("this operation sign is unsupported. got: '%v', operation")
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
		/*
		  numbers[0] = Calc(numbers[0], numbers[1], operators[0]);
		  numbers.RemoveAt(1);
		  operators.RemoveAt(0);
		*/
		numbers[0], err = calculate(numbers[0], numbers[1], operators[0])
		if err != nil {
			return 0, err
		}
		numbers = append(numbers[:1], numbers[2:]...)
		operators = operators[1:]
	}
	return numbers[0], nil
}

/*
	1+(5+(2+1)*3)+7

	1+(14)+7
	15+7
	22
*/

func simplifyParentheses(inp string) (string, error) {
	parOpenCount, parCloseCount := 0, 0
	parFirstOpen, parLastClose := 0, 0
	var res string
	//readyToRead := false
	for i := 0; i < len(inp); i++ {
		cur := inp[i]
		// Если текущий != скобка
		if cur != '(' && cur != ')' {
			// то переписываем как есть
			res += string(cur)
		} else if cur == '(' {
			parOpenCount++
			parFirstOpen = i
		} else if cur == ')' {
			parCloseCount++
			if parOpenCount == parCloseCount {
				parLastClose = i
				temp, err := simplifyParentheses(inp[parFirstOpen : parLastClose+1])
				if err != nil {
					return "", ErrInvalidExpression
				}
				fmt.Println(temp)
			}
		}

		if cur == ')' && parOpenCount == 0 {
			return "", ErrMismatchedParentheses
		}
		if cur == '(' {
			parOpenCount++
			//readyToRead = true
			continue
		} else if cur == ')' {

		}
	}
	return "123", nil
}

func Calc(expression string) (float64, error) {
	//// Проверим что выражение не пустое
	//expression = simplifySpaces(expression)
	//expression = simplifyParentheses(expression)

	/*
		1+(5+(3)*3)+7

		1+(14)+7
		15+7
		22
	*/

	l := len(expression)
	for i := 0; i < l; i++ {
		cur := expression[i] // Текущий символ
		// Если это число то
		if unicode.IsDigit(rune(cur)) {

		}
	}

	//Проверим корректность скобок
	//if checkParentheses(expression) == false {
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

	return 0, nil
}

func main() {
	fmt.Println(calculateSimple("2+2-5+6*34/10-20+5.59"))
	//fmt.Println(strings.FieldsFunc("2+2+2-4+2", separatorOperator))

	//str := "1+(5+(2+1)*3)+7" // 22

	//fmt.Println(str[6:9]) //2+1
	// [i:j] --> i вкл, j не вкл

	//fmt.Println(Calc(str))
}
