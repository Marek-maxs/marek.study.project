package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
)

// 定义 flag cpuprofile
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write mem profile to `file`")

func main() {
	//println(testing.AllocsPerRun(5000, func() {
	//	Tcalcelate()
	//}))
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal().Err(err).Msg("could not create CPU profile: ")
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal().Err(err).Msg("could not start CPU profile: ")
		}
		defer pprof.StopCPUProfile()
	}

	//var wg sync.WaitGroup
	//wg.Add(3000)

	for i := 0; i < 3000; i++ {
		Tcalcelate()
	}

	//wg.Wait()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal().Err(err).Msg("could not create memory profile: ")
		}
		defer f.Close()
		runtime.GC()

		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal().Err(err).Msg("cound not write memory profile: ")
		}
	}
}

func Tcalcelate() {
	// A = B + C - (D+20) - E * （G + H） + R
	str := "{id_507} + {id_508} - ({id_509} + 20) - {id_510} * ({id_511} + {id_512}) + {id_513}"
	//str = strings.ReplaceAll(str, "1*", "")
	//str = strings.ReplaceAll(str, " + -", "-")
	reg := regexp.MustCompile("\\s+")
	expression := reg.ReplaceAllString(str, "")
	expression = strings.ReplaceAll(expression, "1*{", "{")
	expression = strings.ReplaceAll(expression, "+-", "-")
	var datapointIDs []string
	datapointArr := GetExpressionDatapointID(expression)
	for _, strID := range datapointArr {
		datapointIDs = append(datapointIDs, strID[1])
	}
	for _, idValue := range datapointIDs {
		expression = strings.ReplaceAll(expression, fmt.Sprintf("{%s}", idValue), "100")
	}
	resutl, _ := Calculate(expression)
	//fmt.Println(expression)
	fmt.Println(resutl)
	//fmt.Println(err)

	//wg.Done()
}

func GetExpressionDatapointID(expression string) [][]string {
	rex := regexp.MustCompile(`\{(.*?)\}`)
	expressionArr := rex.FindAllStringSubmatch(expression, -1)
	return expressionArr
}


const (
	StackSize = 1000
	BitSize64 = 64
	FortySix  = 46
	ZONE      = 0
	ONE       = 1
	TWO       = 2
	THREE     = 3
	FIVE      = 5
	SEVEN     = 7
	EIGHT     = 8
)

type Stack struct {
	size int
	top  int
	data []interface{}
}

func NewStack(size int) *Stack {
	return &Stack{
		size: size,
		top:  ZONE,
		data: make([]interface{}, size),
	}
}

func (vdc *Stack) IsFull() bool {
	return vdc.top == vdc.size
}

func (vdc *Stack) IsEmpty() bool {
	return vdc.top == ZONE
}

func (vdc *Stack) Pop() (interface{}, bool) {
	if vdc.IsEmpty() {
		return ZONE, false
	}
	vdc.top--
	return vdc.data[vdc.top], true
}

func (vdc *Stack) Push(d interface{}) bool {
	if vdc.IsFull() {
		return false
	}
	vdc.data[vdc.top] = d
	vdc.top++
	return true
}

func (vdc *Stack) Peek() (interface{}, bool) {
	if vdc.IsEmpty() {
		return ZONE, false
	}
	return vdc.data[vdc.top-ONE], true
}

func (vdc *Stack) Len() int {
	return vdc.top
}

func Calculate(expression string) (float64, error) {
	var numStack = NewStack(StackSize)
	var symbolStack = NewStack(StackSize)
	err := expressionHandler(expression, numStack, symbolStack)
	if err != nil {
		return ZONE, err
	}
	for symbolStack.Len() != ZONE {
		if numStack.Len() < TWO {
			res, _ := numStack.Peek()
			return res.(float64), nil
			return 0, errors.New("[7]error in input expression")
		}
		err := operandHandler(numStack, symbolStack)
		if err != nil {
			return ZONE, err
		}
	}
	if numStack.Len() != ONE {
		res, _ := numStack.Peek()
		fmt.Println(numStack.Len())
		fmt.Println(res.(float64))
		return ZONE, errors.New("[10]error in input expression")
	}
	res, _ := numStack.Peek()
	return res.(float64), nil
}

func getSum(data1 float64, data2 float64, c string) (float64, error) {
	switch c {
	case "+":
		return data1 + data2, nil
	case "-":
		return data2 - data1, nil
	case "*":
		return data1 * data2, nil
	case "/":
		return data2 / data1, nil
	}
	return ZONE, errors.New("[11]not a legal symbol")
}

func getPriority(c string) int {
	switch c {
	case "+":
		return THREE
	case "-":
		return TWO
	case "*":
		return FIVE
	case "/":
		return FIVE
	case "(":
		return EIGHT
	case ")":
		return ONE
	default:
		return ZONE
	}
}

func Int32ToFloat64(a int32) (float64, error) {
	b := string(a)
	c, err := strconv.ParseFloat(b, BitSize64)
	return c, err
}

func expressionHandler(expression string, numStack, symbolStack *Stack) error {
	last := -1
	for _, s := range expression {
		p := getPriority(string(s))
		if p == ZONE {
			numberHandler(s, last, numStack, symbolStack)
			last = p
		} else {
			last = p
			if symbolStack.IsEmpty() {
				symbolStack.Push(string(s))
				continue
			}
			t, ok := symbolStack.Peek()
			if !ok {
				return errors.New("[1]the symbol stack is empty and cannot be pushed out")
			}
			top := getPriority(t.(string))
			if (p != ONE && top == EIGHT) || (p != TWO && top == SEVEN) {
				symbolStack.Push(string(s))
				continue
			}
			top, err := symbolHandler(p, top, numStack, symbolStack)
			if err != nil {
				return err
			}
			if p == ONE && top == EIGHT {
				symbolStack.Pop()
				continue
			}
			if p > top {
				symbolStack.Push(string(s))
			} else {
				err := operandHandler(numStack, symbolStack)
				symbolStack.Push(string(s))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func operandHandler(numStack, symbolStack *Stack) error {
	symbol, err := symbolStack.Pop()
	if !err {
		return errors.New("[8] the symbol stack is empty")
	}
	data1, _ := numStack.Pop()
	data2, err := numStack.Pop()
	if !err {
		// modify there 2
		numStack.Push(data1)
		return nil
		// return errors.New("[3] error in input expression")
	}
	res, e := getSum(SwitchFloatType(data1), SwitchFloatType(data2), symbol.(string))
	if e != nil {
		return e
	}
	numStack.Push(res)
	return nil
}

func SwitchFloatType(value interface{}) float64 {
	switch value.(type) {
	case string:
		str := value.(string)
		number, _ := strconv.Atoi(str)
		return float64(number)
	case float64:
		return value.(float64)
	default:
		return 0
	}
	return 0
}

func numberHandler(s int32, last int, numStack, symbolStack *Stack) {
	c, _ := Int32ToFloat64(s)
	if last == ZONE {
		n, _ := numStack.Pop()
		ns := fmt.Sprintf("%v%v", n, c)
		if s == FortySix {
			ns = fmt.Sprintf("%v%v", n, ".")
		}
		if c > 0 {
			c, _ = strconv.ParseFloat(ns, BitSize64)
		} else {
			numStack.Push(ns)
			return
		}
	}
	if last == TWO {
		c = -c
		symbolStack.Pop()
		symbolStack.Push("+")
	}
	numStack.Push(c)
}

func symbolHandler(p, top int, numStack, symbolStack *Stack) (int, error) {
	for p == ONE && top != EIGHT {
		err := operandHandler(numStack, symbolStack)
		if err != nil {
			return ZONE, err
		}
		t, errs := symbolStack.Peek()
		if !errs {
			break
		}
		top = getPriority(t.(string))
	}
	return top, nil
}