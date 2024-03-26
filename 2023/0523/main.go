package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"strconv"
	"testing"
)

type NodeValue struct {
	VirtualDatapointID string
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
}

func parserExpression() {
	// expression := "0.85*{id_5812}+{id_5848}/65580+{id_5849}/32790-{id_5774}-{id_5604}-{id_5615}-{id_5607}-{id_5613}-{id_5612}-{id_5603}-{id_5621}-{id_5645}-{id_5632}-{id_5634}-{id_5637}-{id_5620}-{id_5830}-{id_5642}-{id_5643}-{id_5644}-{id_5633}-{id_5636}-{id_5611}-{id_5616}-{id_5618}-{id_5619}-{id_5624}-{id_5623}-{id_5811}-{id_5628}-{id_5629}-{id_5639}-{id_5640}-{id_5985}-{id_5765}/65580"
	expres := "({id_2007}+{id_2005}+{id_2006}+{id_2008}+({id_2004}+{id_2003})*0.0115+{id_2002}-{id_4213})*0.119+({id_55245}+{id_55246})+0.118+({id_5848}/{id_3322999})"
	// 重构公式
	// calculate := 0
	var tempNode, lastOperator string
	var newNodeValue NodeValue
	masterVirtualDatapointID, _ := uuid.NewV4()
	newNodeValueArr := make([]NodeValue, 0)
	levelCalculateArr := make([]string,len(expres)/2)
	//prevNodeValue := ""
	levelCalculateNum := 0
	cachePrevValue := ""
	// var prevNodeValue int32
	for _, asciiNum := range expres{
		newNodeValue.VirtualDatapointID = masterVirtualDatapointID.String()
		str := string(asciiNum)
		// 正括号
		if asciiNum == SquareBrackets {
			levelCalculateNum, newNodeValueArr = SquareBracketsParser(levelCalculateNum,
				lastOperator, newNodeValue, newNodeValueArr, levelCalculateArr)
			newNodeValue = NodeValue{}
			continue
		}

		isMultipleSign := asciiNum == MultipleSign
		isDivisionSign := asciiNum == DivisionSign
		isPlus := asciiNum == Plus
		isMinusSign := asciiNum == MinusSign
		isAntiCurlyBraces := asciiNum == AntiCurlyBraces

		if isMultipleSign || isDivisionSign || isPlus || isMinusSign || isAntiCurlyBraces {
			var isContinue bool
			isContinue, tempNode, lastOperator, cachePrevValue, newNodeValue,
				newNodeValueArr = parserOperator(tempNode, str, lastOperator, cachePrevValue,
				asciiNum, levelCalculateNum, newNodeValue, newNodeValueArr, levelCalculateArr)
			if isContinue {
				continue
			}
		}
		// 缓存中上一个元素
		if asciiNum == ReverseParentheses || asciiNum == AntiCurlyBraces {
			cachePrevValue = str
		}
		// 反括号
		if asciiNum == ReverseParentheses {
			levelCalculateNum -= 1
			continue
		}

		tempNode += str
	}

	// 如果最后一个值不为空，则意着它是一个常量
	if tempNode != "" {
		tempNodeParser(tempNode, cachePrevValue, lastOperator, newNodeValueArr)
	}

	// 把字符串解析成为新结构的 struct
	// fmt.Println(newNodeValueArr)
}

func parserOperator(tempNode, str, lastOperator, cachePrevValue string, asciiNum int32,
	levelCalculateNum int, newNodeValue NodeValue, newNodeValueArr []NodeValue,
	levelCalculateArr []string) (bool, string, string, string, NodeValue, []NodeValue) {
	// 反 大括号 34
	if tempNode == "" && asciiNum != AntiCurlyBraces {
		lastOperator = str
		newNodeValue.Operator = str

		return true, tempNode, lastOperator, cachePrevValue, newNodeValue, newNodeValueArr
	}

	constant, _ := strconv.ParseFloat(tempNode, 64)
	// 如果是一个常量
	if constant != 0 {
		if cachePrevValue == ")" {
			// 降序，对 常量的计算
			newNodeValueArr = parserConstantAddValue(constant, levelCalculateNum,
				levelCalculateArr, newNodeValue, newNodeValueArr)
			newNodeValue = NodeValue{}
		} else {
			// 如果这个常量前面有一个节点，则需求把常量加到前面的节点中
			newNodeValue.Constant = constant
			newNodeValue.ConstantOperator = str
		}
		tempNode = ""

		return true, tempNode, lastOperator, cachePrevValue, newNodeValue, newNodeValueArr
	}
	// 反 大括号
	if asciiNum == AntiCurlyBraces {
		tempNode += str
		newNodeValueArr = AntiCurlyBracesParser(levelCalculateNum, tempNode, lastOperator,
			levelCalculateArr, newNodeValue, newNodeValueArr)
		newNodeValue = NodeValue{}
		tempNode = ""
		// 缓存最后的元素
		cachePrevValue = str

		return true, tempNode, lastOperator, cachePrevValue, newNodeValue, newNodeValueArr
	}

	return false, tempNode, lastOperator, cachePrevValue, newNodeValue, newNodeValueArr
}

func AntiCurlyBracesParser(levelCalculateNum int, tempNode, lastOperator string, levelCalculateArr []string,
	newNodeValue NodeValue, newNodeValueArr []NodeValue) []NodeValue {
	// 如果是中间节点
	if levelCalculateNum > 0 {
		// 把 virtualDatapointID 换成 中间节点的的ID
		newNodeValue.VirtualDatapointID = levelCalculateArr[levelCalculateNum]
	}
	// 如果当前 元素中没有 操作符 且上一个 操作符不为空，则用上一个操作符
	if lastOperator != "" && newNodeValue.Operator == "" {
		newNodeValue.Operator = lastOperator
	}
	// 关系 节点
	newNodeValue.DependID = tempNode
	newNodeValueArr = append(newNodeValueArr, newNodeValue)

	return newNodeValueArr
}

func SquareBracketsParser(levelCalculateNum int, lastOperator string, newNodeValue NodeValue,
	newNodeValueArr []NodeValue, levelCalculateArr []string) (int, []NodeValue) {
	// 运行级别加1
	levelCalculateNum += 1
	newNodeValue.IsChildNode = 1
	// 本级的 virtual datapoint id
	newVirtualDatapointID, _ := uuid.NewV4()
	levelCalculateArr[levelCalculateNum] = newVirtualDatapointID.String()

	if levelCalculateNum > 1 {
		// 如果计算级别到了1级以上的，这里的virtual datapoint id 是它的上一级
		newNodeValue.VirtualDatapointID = levelCalculateArr[levelCalculateNum - 1]
	}

	newNodeValue.Operator = lastOperator
	newNodeValue.DependID = newVirtualDatapointID.String()
	// 把解析好的数据 压入 数组
	newNodeValueArr = append(newNodeValueArr, newNodeValue)

	return levelCalculateNum, newNodeValueArr
}

func parserConstantAddValue(constant float64, levelCalculateNum int,
	levelCalculateArr []string, newNodeValue NodeValue, newNodeValueArr []NodeValue) []NodeValue {
	for kk := len(newNodeValueArr) -1; kk >= 0; kk-- {
		if newNodeValueArr[kk].IsChildNode == 1 && newNodeValueArr[kk].ConstantOperator == "" &&
			levelCalculateArr[levelCalculateNum+1] == newNodeValueArr[kk].DependID{
			newNodeValueArr[kk].Constant = constant
			newNodeValueArr[kk].ConstantOperator = newNodeValue.Operator
			break
		}
	}

	return newNodeValueArr
}

func tempNodeParser(tempNode, cachePrevValue, lastOperator string, newNodeValueArr []NodeValue) {
	constant, _ := strconv.ParseFloat(tempNode, 64)
	for kk := len(newNodeValueArr)-1; kk >=0 ; kk-- {
		isChildNode := newNodeValueArr[kk].IsChildNode
		constantOperator := newNodeValueArr[kk].ConstantOperator
		if cachePrevValue == ")" && isChildNode == 1 && constantOperator == ""{
			newNodeValueArr[kk].Constant = constant
			newNodeValueArr[kk].ConstantOperator = lastOperator
			break
		} else if cachePrevValue == "}" && constantOperator == "" {
			newNodeValueArr[kk].Constant = constant
			newNodeValueArr[kk].ConstantOperator = lastOperator
			break
		}
	}
}