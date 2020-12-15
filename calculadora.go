package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calc define the strucs
type Calc struct {
	intA     int
	intB     int
	operator string
}

// Operate does the calc
func (Calc) Operate(entrada string, operator string) int {
	entradaLimpia := strings.Split(entrada, operator)
	operator1 := parser(entradaLimpia[0])
	operator2 := parser(entradaLimpia[1])

	total := 0

	switch operator {
	case "+":
		total = operator1 + operator2
		break
	case "-":
		total = operator1 - operator2
		break
	case "*":
		total = operator1 * operator2
		break
	case "/":
		total = operator1 / operator2
		break
	default:
		fmt.Println("Invalid input")
		total = 0
		break
	}

	return total
}

func parser(entrada string) int {
	operator, _ := strconv.Atoi(entrada)
	return operator
}

// ReadInput Receive an input
func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
