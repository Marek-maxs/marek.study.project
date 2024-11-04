package main

import (
	"fmt"
	"math"
	"strings"
)

type formula struct {
}

func (f *formula) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (f *formula) parseNum(s string) (float64, error) {
	return math.ParseFloat(s, 64)
}

func (f *formula) isOperator(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func (f *formula) shouldPop(op1 byte, op2 byte) bool {
	if op1 == '*' || op1 == '/' {
		return true
	}

	return op2 == '+' || op2 == '-'
}

func (f *formula) applyOperator(a float64, b float64, op byte) float64 {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	default:
		panic("unsupported operator")
	}
}

func runFormula() {
	fmt.Println("请输入公式：")
	var formulas string
	fmt.Scanln(&formulas)
	fma := formula{}

	// 将公式中的空格全部去掉
	formulas = strings.ReplaceAll(formulas, " ", "")

	result, err := fma.calculate(formulas)
	if err != nil {
		fmt.Println("计算失败：", err)
		return
	}

	fmt.Println("计算结果：", result)
}

func (f *formula) calculate(formula string) (float64, error) {
	stack := make([]float64, 0) // 使用栈来计算
	ops := make([]byte, 0)

	for i := 0; i < len(formula); i++ {
		if f.isDigit(formula[i]) { // 数字
			j := i
			for ; j < len(formula); j++ {
				if !f.isDigit(formula[j]) && formula[j] != '.' {
					break
				}
			}

			num, _ := f.parseNum(formula[i:j])
			stack = append(stack, num)
			i = j - 1
		} else if f.isOperator(formula[i]) { // 运算符
			for len(ops) > 0 && f.shouldPop(ops[len(ops)-1], formula[i]) {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				nextToTop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				res := f.applyOperator(nextToTop, top, ops[len(ops)-1])

				stack = append(stack, res)
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, formula[i])
		} else {
			return 0, fmt.Errorf("invalid character: %c", formula[i])
		}

		for len(ops) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			nextToTop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res := f.applyOperator(nextToTop, top, ops[len(ops)-1])
			stack = append(stack, res)
			ops = ops[:len(ops)-1]
		}
	}
}
