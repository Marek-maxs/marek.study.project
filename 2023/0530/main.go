package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"strconv"
	"strings"
	"testing"
)

type NodeValue struct {
	VirtualDatapointID string
	Sort int
	Constant float64
	ConstantOperator string
	DependID string
	Operator string
	IsChildNode int
}

const (
	// 正括号
	SquareBrackets = 40
	// 反括号
	ReverseParentheses = 41
	// 反大括号
	AntiCurlyBraces = 125
	// 加号
	Plus = 43
	// 减号
	MinusSign = 45
	// 除号
	DivisionSign = 47
	// 乘号
	MultipleSign = 42
)

var OperatorMap = map[int]string{
	Plus : "+",
	MinusSign: "-",
	DivisionSign: "/",
	MultipleSign: "*",
	AntiCurlyBraces: "}",
}

// Example usage:
func main(){
	fmt.Println(testing.AllocsPerRun(1, func() {
		parserExpression()
	}))
}

type Expression struct {
	LevelCalculateNum int
	TempNode string
	LastOperator string
	CachePrevValue string

	NewNodeValue NodeValue
	NewNodeValueArr []NodeValue
	LevelCalculateArr []string
}

func parserExpression() {
	// expression := "0.85*{id_5812}+{id_5848}/65580+{id_5849}/32790-{id_5774}-{id_5604}-{id_5615}-{id_5607}-{id_5613}-{id_5612}-{id_5603}-{id_5621}-{id_5645}-{id_5632}-{id_5634}-{id_5637}-{id_5620}-{id_5830}-{id_5642}-{id_5643}-{id_5644}-{id_5633}-{id_5636}-{id_5611}-{id_5616}-{id_5618}-{id_5619}-{id_5624}-{id_5623}-{id_5811}-{id_5628}-{id_5629}-{id_5639}-{id_5640}-{id_5985}-{id_5765}/65580"
	expres := "0.5*({id_675488}+{id_33815}+{id_675485})-({id_33821}*2)"
	// 重构公式
	// calculate := 0
	//var tempNode, lastOperator string
	//var newNodeValue NodeValue
	masterVirtualDatapointID, _ := uuid.NewV4()
	//newNodeValueArr := make([]NodeValue, 0)
	//levelCalculateArr := make([]string,len(expres)/2)
	//prevNodeValue := ""
	//levelCalculateNum := 0
	//cachePrevValue := ""
	// var prevNodeValue int32
	newExpression := Expression{
		LevelCalculateNum: 0,
		TempNode:          "",
		LastOperator:      "",
		CachePrevValue:    "",
		NewNodeValue:      NodeValue{},
		NewNodeValueArr:   make([]NodeValue, 0),
		LevelCalculateArr: make([]string,len(expres)/2),
	}
	sort := 0
	for _, asciiNum := range expres{
		sort++
		newExpression.NewNodeValue.VirtualDatapointID = masterVirtualDatapointID.String()
		newExpression.NewNodeValue.Sort = sort
		str := string(asciiNum)
		// 正括号
		if asciiNum == SquareBrackets {
			newExpression.SquareBracketsParser()
			newExpression.NewNodeValue = NodeValue{}
			continue
		}

		if OperatorMap[int(asciiNum)] != "" {
			isContinue := newExpression.parserOperator(str, asciiNum)
			if isContinue {
				continue
			}
		}
		// 缓存中上一个元素
		if asciiNum == ReverseParentheses || asciiNum == AntiCurlyBraces {
			newExpression.CachePrevValue = str
		}
		// 反括号
		if asciiNum == ReverseParentheses {
			newExpression.LevelCalculateNum -= 1
			newExpression.CachePrevValue = str
			continue
		}

		newExpression.TempNode += str
	}

	// 如果最后一个值不为空，则意着它是一个常量
	if newExpression.TempNode != "" {
		newExpression.tempNodeParser()
	}

	// 把字符串解析成为新结构的 struct
	fmt.Println(newExpression.NewNodeValueArr)
}

func (e *Expression)parserOperator(str string, asciiNum int32,
	) (bool) {
	// 反 大括号 34
	if e.TempNode == "" && asciiNum != AntiCurlyBraces {
		e.LastOperator = str
		e.NewNodeValue.Operator = str

		return true
	}

	constant, _ := strconv.ParseFloat(e.TempNode, 64)
	// 如果是一个常量
	if constant != 0 {
		if e.CachePrevValue == ")" {
			if OperatorMap[int(asciiNum)] != "" {
				e.LastOperator = str
			}
			// 降序，对 常量的计算
			e.parserConstantAddValue(constant)
			e.NewNodeValue = NodeValue{}
		} else {
			// 如果这个常量前面有一个节点，则需求把常量加到前面的节点中
			e.NewNodeValue.Constant = constant
			e.NewNodeValue.ConstantOperator = str
		}
		e.TempNode = ""

		return true
	}
	// 反 大括号
	if asciiNum == AntiCurlyBraces {
		refIDStr := e.TempNode
		refID := strings.Replace(refIDStr, "{id_", "", 1)
		fmt.Println(refID)
		e.TempNode += str
		e.AntiCurlyBracesParser()
		e.NewNodeValue = NodeValue{}
		e.TempNode = ""
		// 缓存最后的元素
		e.CachePrevValue = str

		return true
	}

	return false
}

func (e *Expression)AntiCurlyBracesParser() {
	// 如果是中间节点
	if e.LevelCalculateNum > 0 {
		// 把 virtualDatapointID 换成 中间节点的的ID
		e.NewNodeValue.VirtualDatapointID = e.LevelCalculateArr[e.LevelCalculateNum]
	}
	// 如果当前 元素中没有 操作符 且上一个 操作符不为空，则用上一个操作符
	if e.LastOperator != "" && e.NewNodeValue.Operator == "" && e.CachePrevValue != ")" {
		tmp := e.NewNodeValue.Operator
		e.NewNodeValue.Operator = e.LastOperator
		e.LastOperator = tmp
	}
	// 关系 节点
	e.NewNodeValue.DependID = e.TempNode
	e.NewNodeValueArr = append(e.NewNodeValueArr, e.NewNodeValue)
}

func (e *Expression)SquareBracketsParser() {
	// 运行级别加1
	e.LevelCalculateNum += 1
	e.NewNodeValue.IsChildNode = 1
	// 本级的 virtual datapoint id
	newVirtualDatapointID, _ := uuid.NewV4()
	e.LevelCalculateArr[e.LevelCalculateNum] = newVirtualDatapointID.String()

	if e.LevelCalculateNum > 1 {
		// 如果计算级别到了1级以上的，这里的virtual datapoint id 是它的上一级
		e.NewNodeValue.VirtualDatapointID = e.LevelCalculateArr[e.LevelCalculateNum - 1]
	}

	e.NewNodeValue.Operator = e.LastOperator
	e.NewNodeValue.DependID = newVirtualDatapointID.String()
	// 把解析好的数据 压入 数组
	e.NewNodeValueArr = append(e.NewNodeValueArr, e.NewNodeValue)
}

func (e *Expression)parserConstantAddValue(constant float64) {
	for kk := len(e.NewNodeValueArr) -1; kk >= 0; kk-- {
		if e.NewNodeValueArr[kk].IsChildNode == 1 && e.NewNodeValueArr[kk].ConstantOperator == "" &&
			e.LevelCalculateArr[e.LevelCalculateNum+1] == e.NewNodeValueArr[kk].DependID{
			e.NewNodeValueArr[kk].Constant = constant
			e.NewNodeValueArr[kk].ConstantOperator = e.NewNodeValue.Operator
			break
		}
	}
}

func (e *Expression)tempNodeParser() {
	constant, _ := strconv.ParseFloat(e.TempNode, 64)
	for kk := len(e.NewNodeValueArr)-1; kk >=0 ; kk-- {
		isChildNode := e.NewNodeValueArr[kk].IsChildNode
		constantOperator := e.NewNodeValueArr[kk].ConstantOperator
		if e.CachePrevValue == ")" && isChildNode == 1 && constantOperator == ""{
			e.NewNodeValueArr[kk].Constant = constant
			e.NewNodeValueArr[kk].ConstantOperator = e.LastOperator
			break
		} else if e.CachePrevValue == "}" && constantOperator == "" {
			e.NewNodeValueArr[kk].Constant = constant
			e.NewNodeValueArr[kk].ConstantOperator = e.LastOperator
			break
		}
	}
}