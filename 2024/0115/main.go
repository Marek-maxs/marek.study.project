package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"
	"strings"
)

/**
*
* Author: Marek
* Date: 2024-01-15 12:16
* Email: 364021318@qq.com
*
 */

// 定义 flag cpuprofile
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write mem profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Println("could not create CPU profile: ")
			log.Fatal(err)
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Println("could not start CPU profile: ")
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	//var wg sync.WaitGroup
	//wg.Add(3000)

	for i := 0; i < 3000; i++ {
		go Tcalcelate()
	}

	//wg.Wait()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			fmt.Println("could not create memory profile: ")
			log.Fatal(err)
		}
		defer f.Close()
		runtime.GC()

		if err := pprof.WriteHeapProfile(f); err != nil {
			fmt.Println("cound not write memory profile: ")
			log.Fatal(err)
		}
	}
}

func Tcalcelate() {
	arr := map[string]float64{"507":100,"508":100,"509":100,"510":100,"511":100,"512":100,"513":100}
	str := "{id_507} + {id_508} - ({id_509} + 20) - {id_510} * ({id_511} + {id_512}) + {id_513}"

	for k, val := range arr {
		str = strings.ReplaceAll(str, fmt.Sprintf("{id_%s}", k), fmt.Sprintf("%f", val))
	}

	newExpress := "print('{}'.format(eval('%s')))"
	pycodeGo := fmt.Sprintf(newExpress, str)
	cmd := exec.Command("python", "-c", newExpress) // 要执行的 Python 代码
	output, err := cmd.Output()                                       // 获取输出结果
	if err != nil {
		log.Fatal(err)
	}
	result := string(output)
	fmt.Println(result)
	//wg.Done()
}

func oldFunction() {
	express := `
import re

# 输入字符串
string = "%s"
# 使用split()函数按'&'进行分割
parts = string.split('&')
 
# 定义空数组
items = {}
# 利用列表生成式获取每个键值对并添加到结果数组中
for part in parts:
	if '=' in part:
		p = part.split('=')
		items[p[0]] = p[1]

print(items)
# items = {5696: 100, 5697: 100, 5698: 100, 5699: 100, 5700: 100, 5701: 100, 5702: 100, 5703: 100, 5704: 100}
_component_values = {"id_"+str(k): items.get(k) if items.get(k) is not None else 0 for k in items}
print(_component_values)
composition_expression = "{id_5696}-{id_5697}-{id_5698}-{id_5699}-{id_5700}-{id_5701}-{id_5702}-{id_5703}-{id_5704}"
systemIDs = re.findall("id_[0-9]+", composition_expression)            
for sysID in systemIDs:                    
    if not _component_values.get(sysID):
        # print("_component_values[{}] = 0".format(sysID))
        _component_values[sysID] = 0

print("updated _component_values:{}".format(_component_values))
expression = composition_expression.format(**_component_values)
expression=expression.replace('None', "0")
if not re.match(r'^[0-9\+\-\*/e\. \(\)]*$', expression):
    print("Invalid expression, potentially dangerous to call EVAL on", expression)
    
print(expression)
value = eval(expression)
print("get_composed_value:{}".format(value))
`

	refIDValue := "5696=900&5697=100&5698=100&5699=100&5700=100&5701=100&5702=100&5703=100&5704=100"
	express = fmt.Sprintf(express, refIDValue)
	cmd := exec.Command("python", "-c", express) // 要执行的 Python 代码
	output, err := cmd.Output()                                       // 获取输出结果
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}