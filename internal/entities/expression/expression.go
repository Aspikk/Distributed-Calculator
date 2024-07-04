package expression

import (
	"strings"
	"unicode"

	stack "github.com/Aspikk/Distributed-Calculator/internal/entities/stack"
)

var (
	id = 1
)

type Expression struct {
	ID     int     `json:"id"`
	Result float64 `json:"result"`
	Status string  `json:"status"`
	Raw    string  `json:"expression"`
	rpn    string  // Reverse Polish Notation
}

func New() *Expression {
	return &Expression{
		ID:     id,
		Result: 0,
		Status: "new",
	}
}

// Methods

func (e *Expression) AddSpaces() {
	res := ""
	for _, v := range e.Raw {
		if v != ' ' {
			res += string(v)
		}
	}

	e.Raw = ""

	for i, v := range res {
		if unicode.IsDigit(v) {
			e.Raw += string(v)
		} else {
			idx := i + 1
			if i == len(res)-1 {
				idx = i
			}
			if unicode.IsDigit(rune(res[idx])) {
				e.Raw += " " + string(v) + " "
			} else {
				e.Raw += " " + string(v)
			}
		}
	}

	e.Raw = strings.Trim(e.Raw, " ")
}

func (e *Expression) SetStatus(status string) {
	e.Status = status
}

func (e *Expression) AddTo(storage *[]*Expression) {
	*storage = append(*storage, e)
}

func (e *Expression) IsInvalid() bool {
	if e.Raw == "" {
		return true
	}
	if OnlyParentheses(e.Raw) {
		return true
	}

	stack := stack.NewStack[rune]()

	operationCount := 0

	for i, v := range e.Raw {
		if v == ' ' {
			continue
		}

		if !(isStringDigit(string(v)) || isOperation(v) || v == '(' || v == ')') {
			return true
		}

		if isOperation(v) {
			if i == 0 {
				return true
			}

			if (i > 0) && (e.Raw[i-1] == '(' || isOperation(rune(e.Raw[i-1]))) {
				return true
			}

			operationCount++
		}

		if v == '(' {
			if i >= 2 && (!isOperation(rune(e.Raw[i-2])) && e.Raw[i-2] != '(') {
				return true
			}
			if i+2 < len(e.Raw) && isOperation(rune(e.Raw[i+2])) {
				return true
			}
			stack.Push(v)
		}

		if v == ')' {
			if !isStringDigit(string(e.Raw[i-2])) {
				return true
			}
			if i+2 < len(e.Raw) && isStringDigit(string(e.Raw[i+2])) {
				return true
			}
			_, ok := stack.Pop()
			if !ok {
				return true
			}
		}
	}

	lastChar := e.Raw[len(e.Raw)-1]
	if isOperation(rune(lastChar)) {
		return true
	}

	if operationCount < 1 {
		return true
	}

	if !stack.IsEmpty() {
		return true
	}

	id++

	return false
}

func (e *Expression) ToRpn() {
	var priorities map[string]int = map[string]int{
		"^": 4,
		"*": 3,
		"/": 3,
		"+": 2,
		"-": 2,
		"(": 1,
	}

	stack := stack.NewStack[string]()

	for _, v := range strings.Split(e.Raw, " ") {
		if isStringDigit(v) {
			e.rpn += v + " "
		} else {
			if v == "(" {
				stack.Push(v)
				continue
			} else if v == ")" {
				for {
					if stack.GetTop() == "(" {
						stack.Pop()
						break
					}
					val, _ := stack.Pop()
					e.rpn += val + " "
				}
				continue
			}

			if stack.IsEmpty() || priorities[stack.GetTop()] < priorities[v] {
				stack.Push(v)
			} else {
				for val, is := stack.Pop(); priorities[val] >= priorities[v] && is; {
					if val == "(" {
						val, is = stack.Pop()
						continue
					}
					e.rpn += val + " "
					val, is = stack.Pop()
				}
				stack.Push(v)
			}
		}
	}

	for val, is := stack.Pop(); is; {
		if val == ")" {
			val, is = stack.Pop()
			continue
		}

		e.rpn += val + " "
		val, is = stack.Pop()
	}
}

// Help functions

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

func isStringDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func isOperation(v rune) bool {
	return v == '+' || v == '-' || v == '*' || v == '/' || v == '^'
}
