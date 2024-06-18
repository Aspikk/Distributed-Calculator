package entities

import (
	"unicode"

	entities "github.com/Aspikk/Distributed-Calculator/internal/entities/stack"
)

type Expression struct {
	ID  int    `json:"id"`
	Raw string `json:"expression"`
}

func (e *Expression) RemoveSpaces() {
	res := ""
	for _, v := range e.Raw {
		if v != ' ' {
			res += string(v)
		}
	}
	e.Raw = res
}

func (e *Expression) SetID(id *int) {
	e.ID = *id
	*id = *id + 1
}

func (e *Expression) IsInvalid() bool {
	if e.Raw == "" {
		return true
	}

	stack := entities.NewStack[rune]()

	for _, v := range e.Raw {
		if !(unicode.IsDigit(rune(v)) || v == '+' || v == '-' || v == '*' || v == '/' || v == '(' || v == ')' || v == '^') {
			return true
		}

		if v == '(' {
			stack.Push(v)
		}

		if v == ')' {
			_, err := stack.Pop()
			if err {
				return true
			}
		}
	}

	return !stack.IsEmpty()
}
