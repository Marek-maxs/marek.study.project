package main

import (
	"flag"
	"fmt"
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

type NodeData struct {
	RelashData map[string][]NodeValue
	InfluxdbData map[string]float64
	Result float64
}

// 定义 flag cpuprofile
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write mem profile to `file`")
func main()  {
	println(testing.AllocsPerRun(5000, func() {
		calculate()
	}))
	fmt.Println(testing.AllocsPerRun(5000, func() {
		calculate()
	}))
	//flag.Parse()
	//if *cpuprofile != "" {
	//	f, err := os.Create(*cpuprofile)
	//	if err != nil {
	//		log.Fatal().Err(err).Msg("could not create CPU profile: ")
	//	}
	//	defer f.Close()
	//
	//	if err := pprof.StartCPUProfile(f); err != nil {
	//		log.Fatal().Err(err).Msg("could not start CPU profile: ")
	//	}
	//	defer pprof.StopCPUProfile()
	//}
	//
	//var wg sync.WaitGroup
	//wg.Add(5000)
	//
	//for i := 0; i < 5000; i++ {
	//	go calculate()
	//}
	//
	//wg.Wait()
	//
	//if *memprofile != "" {
	//	f, err := os.Create(*memprofile)
	//	if err != nil {
	//		log.Fatal().Err(err).Msg("could not create memory profile: ")
	//	}
	//	defer f.Close()
	//	runtime.GC()
	//
	//	if err := pprof.WriteHeapProfile(f); err != nil {
	//		log.Fatal().Err(err).Msg("cound not write memory profile: ")
	//	}
	//}

}

func calculate() {
	// 从数据中得到的数据
	data := []NodeValue{
		{
			VirtualDatapointID: "A",
			DependID:"B",
			Operator:"",
			IsChildNode:0,
			Constant:0,
			ConstantOperator:"",
		},
		{
			VirtualDatapointID: "A",
			DependID:"C",
			Operator:"+",
			IsChildNode:0,
			Constant:0,
			ConstantOperator:"",
		},
		{
			VirtualDatapointID: "A",
			DependID:"D",
			Operator:"-",
			IsChildNode:0,
			Constant:20,
			ConstantOperator:"+",
		},
		{
			VirtualDatapointID: "A",
			DependID:"X",
			Operator:"-",
			IsChildNode:1,
			Constant:0,
			ConstantOperator:"",
		},
		{
			VirtualDatapointID: "X",
			DependID:"H",
			Operator:"",
			IsChildNode:0,
			Constant:0,
			ConstantOperator:"",
		},
		{
			VirtualDatapointID: "X",
			DependID:"G",
			Operator:"+",
			IsChildNode:0,
			Constant:0,
			ConstantOperator:"",
		},
		{
			VirtualDatapointID: "X",
			DependID:"E",
			Operator:"*",
			IsChildNode:0,
			Constant:0,
			ConstantOperator:"",
		},
		{
			VirtualDatapointID: "A",
			DependID:"R",
			Operator:"+",
			IsChildNode:1,
			Constant:0,
			ConstantOperator:"",
		},
		{
			VirtualDatapointID: "R",
			DependID:"K",
			Operator:"",
			IsChildNode:0,
			Constant:0,
			ConstantOperator:"",
		},
		{
			VirtualDatapointID: "R",
			DependID:"L",
			Operator:"*",
			IsChildNode:0,
			Constant:0,
			ConstantOperator:"",
		},
		{
			VirtualDatapointID: "R",
			DependID:"Z",
			Operator:"+",
			IsChildNode:0,
			Constant:0,
			ConstantOperator:"",
		},

	}

	// 从influxdb 中获取到的数据
	influxdbData := map[string]float64{
		"B": 100,
		"C": 100,
		"D": 10,
		"G": 10,
		"H": 10,
		"E": 10,
		"Z": 23,
		"L": 30,
		"K": 10,
	}

	newRelashData := make(map[string][]NodeValue)
	for index, value := range data{
		if newRelashData[value.VirtualDatapointID] == nil {
			newRelashData[value.VirtualDatapointID] = make([]NodeValue, len(data))
		}
		newRelashData[value.VirtualDatapointID][index] = value
	}

	newDodeData := NodeData{
		RelashData:   newRelashData,
		InfluxdbData: influxdbData,
	}
	// 获取A 虚拟节点的公式
	// A = B + C - (D+20) - E * （G + H） + R
	// A = 100 + 100 - 30 - 10 * (10+10) + (23+30*10)
	result := newDodeData.CalculateVirtualDatapoint("A")
	fmt.Println("结果：")
	fmt.Println(result)
}


// Helper function to perform the calculation for a node
func (g *NodeData)OperatorParser(operator string, result, value float64) float64 {
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
		if value.DependID == "" {
			continue
		}
		// 如果是嵌套虚拟节点的
		if value.IsChildNode == 1 {
			childResult := g.CalculateVirtualDatapoint(value.DependID)
			if value.Constant != 0 {
				childResult = g.OperatorParser(value.ConstantOperator, childResult, value.Constant)
			}
			result = g.OperatorParser(value.Operator, result, childResult)
		} else {
			khwValue := g.InfluxdbData[value.DependID]
			if value.Constant != 0 {
				khwValue = g.OperatorParser(value.ConstantOperator, khwValue, value.Constant)
			}
			result = g.OperatorParser(value.Operator, result, khwValue)
		}
	}

	return result
}