package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) Operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	op1 := parse(entradaLimpia[0])
	op2 := parse(entradaLimpia[1])
	switch operador {
	case "+":
		fmt.Println(op1 + op2)
		return op1 + op2
	case "-":
		fmt.Println(op1 - op2)
		return op1 - op2
	case "*":
		fmt.Println(op1 * op2)
		return op1 * op2
	case "/":
		fmt.Println(op1 / op2)
		return op1 / op2
	default:
		fmt.Println(operador, "no está soportado")
		return 0
	}
}

func Parse(entrada string) int {
	op, _ := strconv.Atoi(entrada)
	return op
}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
