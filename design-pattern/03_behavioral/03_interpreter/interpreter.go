package interpreter

import (
	"strconv"
	"strings"
)

// Expression 表达式接口
type Expression interface {
	Interpret() int
}

// NumberExpression 数字表达式（终结符表达式）
type NumberExpression struct {
	number int
}

func NewNumberExpression(number int) *NumberExpression {
	return &NumberExpression{number: number}
}

func (n *NumberExpression) Interpret() int {
	return n.number
}

// AddExpression 加法表达式（非终结符表达式）
type AddExpression struct {
	left, right Expression
}

func NewAddExpression(left, right Expression) *AddExpression {
	return &AddExpression{
		left:  left,
		right: right,
	}
}

func (a *AddExpression) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

// SubtractExpression 减法表达式（非终结符表达式）
type SubtractExpression struct {
	left, right Expression
}

func NewSubtractExpression(left, right Expression) *SubtractExpression {
	return &SubtractExpression{
		left:  left,
		right: right,
	}
}

func (s *SubtractExpression) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}

// Parser 解析器
type Parser struct {
	tokens []string
	pos    int
}

func NewParser(expression string) *Parser {
	tokens := strings.Fields(expression)
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

func (p *Parser) Parse() Expression {
	token := p.tokens[p.pos]
	p.pos++

	// 如果是数字，创建数字表达式
	if num, err := strconv.Atoi(token); err == nil {
		return NewNumberExpression(num)
	}

	// 如果是操作符，创建相应的表达式
	left := p.Parse()
	right := p.Parse()

	switch token {
	case "+":
		return NewAddExpression(left, right)
	case "-":
		return NewSubtractExpression(left, right)
	default:
		panic("未知的操作符: " + token)
	}
}
