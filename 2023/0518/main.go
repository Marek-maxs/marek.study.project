package main

import (
	"github.com/gofrs/uuid"
	"reflect"
	"strconv"
)

// Define a struct to represent a node in the graph
type Node struct {
	ID int  // virtual datapoint
	Dependencies []Dependency
	Operator string
	Result float64
}

// 真实的datapoint
type Dependency struct {
	NodeID string // datapoint id
	Operator string
	Value float64 // datapoint value
}

// Define a struct to represent the graph
type Graph struct {
	Nodes []*Node  // 整个计算公式
	Result float64 // 最后的计算结果
}

// Function to add n node to the graph
func(g *Graph) AddNode(dependencies []Dependency, operator string) {
	node := &Node{
		ID:           0,
		Dependencies: dependencies,
		Operator: operator,
		Result: 0.0,
	}
	g.Nodes = append(g.Nodes, node)
}

// Function to get the node by ID
func(g *Graph) GetNode(id int) *Node{
	for _, node := range g.Nodes {
		if node.ID == id {
			return node
		}
	}
	return nil
}

type NodeData struct {
	RelashData map[string][]map[string]interface{}
	InfluxdbData map[string]float64
}

type NodeValue struct {
	VirtualDatapointID string
	Constant float64
	ConstantOperator string
	DependID string
	Operator string
	IsChildNode int
}

// Example usage:
func main(){
	// expression := "0.85*{id_5812}+{id_5848}/65580+{id_5849}/32790-{id_5774}-{id_5604}-{id_5615}-{id_5607}-{id_5613}-{id_5612}-{id_5603}-{id_5621}-{id_5645}-{id_5632}-{id_5634}-{id_5637}-{id_5620}-{id_5830}-{id_5642}-{id_5643}-{id_5644}-{id_5633}-{id_5636}-{id_5611}-{id_5616}-{id_5618}-{id_5619}-{id_5624}-{id_5623}-{id_5811}-{id_5628}-{id_5629}-{id_5639}-{id_5640}-{id_5985}-{id_5765}/65580"
	expres := "({id_2007}+{id_2005}+{id_2006}+{id_2008}+({id_2004}+{id_2003})+{id_2002}-{id_4213})*0.119+({id_55245}+{id_55246})+0.118+({id_5848}/{id_3322999})"
	// 重构公式
	// calculate := 0
	var tempNode, lastOperator string
	var newNodeValue NodeValue
	masterVirtualDatapointID, _ := uuid.NewV4()
	newNodeValueArr := make([]NodeValue, 0)
	levelCalculateArr := make([]string,len(expres)/2)
	//prevNodeValue := ""
	levelCalculateNum := 0
	newNodeValueNum := 0
	cachePrevValue := ""
	// var prevNodeValue int32
	for _, v := range expres{
		newNodeValue.VirtualDatapointID = masterVirtualDatapointID.String()
		p := string(v)
		// 正括号
		if v == 40 {
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
			newNodeValueNum += 1
			// 把解析好的数据 压入 数组
			newNodeValueArr = append(newNodeValueArr, newNodeValue)
			newNodeValue = NodeValue{}
			continue
		}

		if p == "*" || p == "/" || p == "+" || p == "-" || p == "}" {
			lastOperator = p

			// 反 大括号 34
			if tempNode == "" && v != 34 {
				newNodeValue.Operator = p
				continue
			}
			c, _ := strconv.ParseFloat(tempNode, 64)
			// 如果是一个常量
			if c != 0 {
				if cachePrevValue == ")" {
					index := 0
					// 降序，对 常量的计算
					for kk := len(newNodeValueArr) -1; kk >= 0; kk-- {
						if newNodeValueArr[kk].IsChildNode == 1 && newNodeValueArr[kk].ConstantOperator == ""{
							index = kk
							break
						}
					}

					newNodeValueArr[index].Constant = c
					newNodeValueArr[index].ConstantOperator = newNodeValue.Operator
					newNodeValue = NodeValue{}
				} else {
					// 如果这个常量前面有一个节点，则需求把常量加到前面的节点中
					newNodeValue.Constant = c
					newNodeValue.ConstantOperator = p
				}
				tempNode = ""
				continue
			}
			// 反 大括号
			if v == 125 {
				tempNode += p
				newNodeValue.DependID = tempNode
				// 如果是中间节点
				if levelCalculateNum > 0 {
					// 把 virtualDatapointID 换成 中间节点的的ID
					newNodeValue.VirtualDatapointID = levelCalculateArr[levelCalculateNum]
					// 如果当前计算级别不是第一级，则需要把它的virtual datapoint id 改为它的上一级
					//if levelCalculateNum != 0 && (levelCalculateNum - 1) != 0 {
					//	newNodeValue.VirtualDatapointID = levelCalculateArr[levelCalculateNum -1]
					//}
				}
				// 关系 节点
				newNodeValueNum += 1
				newNodeValueArr = append(newNodeValueArr, newNodeValue)
				newNodeValue = NodeValue{}
				tempNode = ""
				// 缓存最后的元素
				cachePrevValue = p
				continue
			}
		}
		// 缓存中上一个元素
		if p == ")" || p == "}" {
			cachePrevValue = p
		}
		// 反括号
		if v == 41 {
			levelCalculateNum -= 1
			continue
		}

		tempNode += p
	}

	// 如果最后一个值不为空，则意着它是一个常量
	if tempNode != "" {
		c, _ := strconv.ParseFloat(tempNode, 64)
		for kk := len(newNodeValueArr)-1; kk >=0 ; kk-- {
			isChildNode := newNodeValueArr[kk].IsChildNode
			constantOperator := newNodeValueArr[kk].ConstantOperator
			if cachePrevValue == ")" && isChildNode == 1 && constantOperator == ""{
				newNodeValueArr[kk].Constant = c
				newNodeValueArr[kk].ConstantOperator = lastOperator
				break
			} else if cachePrevValue == "}" && constantOperator == "" {
				newNodeValueArr[kk].Constant = c
				newNodeValueArr[kk].ConstantOperator = lastOperator
				break
			}
		}
	}
}

// Helper function to perform the calculation for a node
func operatorParser(operator string, result, value float64) float64 {
	switch operator {
	case "+":
		result += value
	case "-":
		result -= value
	case "*":
		result *= value
	case "/":
		result /= value
	default:
		result = value
	}

	return result
}

func (g *NodeData)CalculateVirtualDatapoint(virtualDatapointID string) float64 {
	var result float64
	for _, value := range g.RelashData[virtualDatapointID] {
		// 如果是嵌套虚拟节点的
		if value["isChildNode"].(int) == 1 {
			childResult := g.CalculateVirtualDatapoint(value["dependID"].(string))
			if value["constant"].(int) != 0 {
				childResult = operatorParser(value["constant_operator"].(string), childResult, float64(value["constant"].(int)))
			}
			result = operatorParser(value["operator"].(string), result, childResult)
		} else {
			khwValue := g.InfluxdbData[value["dependID"].(string)]
			if value["constant"].(int) != 0 {
				khwValue = operatorParser(value["constant_operator"].(string), khwValue, float64(value["constant"].(int)))
			}
			result = operatorParser(value["operator"].(string), result, khwValue)
		}
	}

	return result
}

func clear(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}