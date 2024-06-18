package entities

import (
	"unicode"

	entities "github.com/Aspikk/Distributed-Calculator/internal/entities/stack"
)

type Expression struct {
	ID  int    `json:"id"`
	Raw string `json:"expression"`
}

func (e *Expression) RemoveSpaces() *Expression {
	res := ""
	for _, v := range e.Raw {
		if v != ' ' {
			res += string(v)
		}
	}
	e.Raw = res

	return e
}

func (e *Expression) SetID(id *int) {
	e.ID = *id
	*id = *id + 1
}

func (e *Expression) IsInvalid() bool {
	if e.Raw == "" {
		return true
	}
	if OnlyParentheses(e.Raw) {
		return true
	}

	stack := entities.NewStack[rune]()

	for i, v := range e.Raw {
		if !(unicode.IsDigit(rune(v)) || v == '+' || v == '-' || v == '*' || v == '/' || v == '(' || v == ')' || v == '^') {
			return true
		}

		if v == '+' || v == '*' || v == '/' || v == '^' || v == '-' {
			if i == 0 {
				return true
			}

			if (i > 0) && (e.Raw[i-1] == '(' || e.Raw[i-1] == '+' || e.Raw[i-1] == '-' || e.Raw[i-1] == '*' || e.Raw[i-1] == '/' || e.Raw[i-1] == '^') {
				return true
			}
		}

		if v == '(' {
			stack.Push(v)
		}

		if v == ')' {
			_, ok := stack.Pop()
			if !ok {
				return true
			}
		}
	}

	lastChar := e.Raw[len(e.Raw)-1]
	if lastChar == '+' || lastChar == '-' || lastChar == '*' || lastChar == '/' || lastChar == '^' {
		return true
	}

	return !stack.IsEmpty()
}

func OnlyParentheses(expression string) bool {
	hasOnlyParenthesis := true

	for _, char := range expression {
		if char != '(' && char != ')' {
			hasOnlyParenthesis = false
			break
		}
	}

	return hasOnlyParenthesis
}
