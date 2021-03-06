package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

type Token interface {
	Interpret() int
}

type Number struct {
	val int
}
func (n *Number) Interpret() int {
	return n.val
}

type Add struct {
	left, right Token
}
func (t *Add) Interpret() int {
	return t.left.Interpret() + t.right.Interpret()
}

type Subtract struct {
	left, right Token
}
func (t *Subtract) Interpret() int {
	return t.left.Interpret() - t.right.Interpret()
}

type Interpreter struct {
	expression []string
	variables  map[string]int
	index      int
	prev       Token
}
func (i *Interpreter) Parse(expression string, variables map[string]int) {
	i.expression = strings.Split(expression, " ")
	i.variables = variables

	for i.index < len(i.expression) {
		switch i.expression[i.index] {
		case "+":
			i.index++
			i.prev = &Add{
				left:  i.prev,
				right: i.parseNumber(),
			}
		case "-":
			i.index++
			i.prev = &Subtract{
				left:  i.prev,
				right: i.parseNumber(),
			}
		default:
			i.prev = i.parseNumber()
		}
	}
}
func (i *Interpreter) parseNumber() Token {
	tokenStr := i.expression[i.index]
	i.index++
	val, err := strconv.Atoi(tokenStr)
	if err == nil {
		// token at index can be parsed as a number
		return &Number{val}
	}
	if val, ok := i.variables[tokenStr]; ok {
		// token at index can be parsed as a known variable
		return &Number{val}
	}
	panic(fmt.Sprintf("Invalid token %s", tokenStr))
}
func (i *Interpreter) Interpret() int {
	return i.prev.Interpret()
}

func main() {
	i := &Interpreter{}
	variables := map[string]int{
		"x": 1,
		"y": 2,
		"z": 3,
	}
	i.Parse("x + y + z + 4 - 5", variables)
	result := i.Interpret()
	fmt.Printf("result: %d\n", result)
}
