package main

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"regexp"
	"strings"
	"time"
)

func main() {
	//src        := "~{;(<dR;:x>F#,6@WCN^O`GW!#"
	src := "¤SF@µ°͟ݴر"
	srcCharset := "ShiftJIS"
	dstCharset := "UTF-8"
	_, err := gcharset.Convert(dstCharset, srcCharset, src)
	if err != nil {
		fmt.Println(err)
	}
	_ = transform.NewReader(strings.NewReader(src), japanese.ShiftJIS.NewDecoder())
	//decBytes, _ := io.ReadAll(r)
	//decS := string(decBytes)
	//fmt.Println(decS)
	//fmt.Println(uint8(0x05))
	//year := time.Now().Year()
	//month := time.Now().Month()
	//fmt.Println(year)
	//fmt.Println(month)
	//fmt.Println(time.Unix(1675511947, 0).Minute() - time.Unix(1675511347, 0).Minute())
	//fmt.Println(time.Monday)
	str := "20230207113001"
	parse, _ := time.Parse("20060201150405", str)
	fmt.Println(parse.String())
	rex := regexp.MustCompile(`\{(.*?)\}`)
	expressionArr := rex.FindAllStringSubmatch("-100 * {id_5794}", -1)
	for _, v := range expressionArr {
		fmt.Println(v[1])
	}
	//fmt.Println(expressionArr)
	//var ap = map[int]int{
	//	1:1,
	//	2:2,
	//}
	//newAp, _ := json.Marshal(ap)
	//fmt.Println(string(newAp))
	//strs := `{"1":1,"2":2}`
	//var aq map[int]int
	//aq = make(map[int]int)
	//json.Unmarshal([]byte(strs), &aq)
	//fmt.Println(aq[2])
	//fmt.Println(fmt.Sprintf("abc %s", errors.New("how are you")))
}