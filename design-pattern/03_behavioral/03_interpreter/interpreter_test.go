package interpreter

import "testing"

func TestInterpreter(t *testing.T) {
	// 测试简单的数字表达式
	numExp := NewNumberExpression(5)
	if result := numExp.Interpret(); result != 5 {
		t.Errorf("数字表达式解释错误，期望5，得到%d", result)
	}

	// 测试加法表达式
	addExp := NewAddExpression(
		NewNumberExpression(5),
		NewNumberExpression(3),
	)
	if result := addExp.Interpret(); result != 8 {
		t.Errorf("加法表达式解释错误，期望8，得到%d", result)
	}

	// 测试减法表达式
	subExp := NewSubtractExpression(
		NewNumberExpression(10),
		NewNumberExpression(4),
	)
	if result := subExp.Interpret(); result != 6 {
		t.Errorf("减法表达式解释错误，期望6，得到%d", result)
	}

	// 测试复杂表达式
	complexExp := NewAddExpression(
		NewSubtractExpression(
			NewNumberExpression(10),
			NewNumberExpression(4),
		),
		NewNumberExpression(5),
	)
	if result := complexExp.Interpret(); result != 11 {
		t.Errorf("复杂表达式解释错误，期望11，得到%d", result)
	}

	// 测试解析器
	parser := NewParser("+ - 10 4 5")
	result := parser.Parse().Interpret()
	if result != 11 {
		t.Errorf("解析器解释错误，期望11，得到%d", result)
	}

	// 测试另一个复杂表达式
	parser = NewParser("- + 5 3 4")
	result = parser.Parse().Interpret()
	if result != 4 {
		t.Errorf("解析器解释错误，期望4，得到%d", result)
	}
}
