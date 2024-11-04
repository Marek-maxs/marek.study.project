package main

import (
	"fmt"
	"strings"
)

type LiveDataTime struct {
	M00 float64 `json:"m00"`
	M01 float64 `json:"m01"`
	M02 float64 `json:"m02"`
	M03 float64 `json:"m03"`
	M04 float64 `json:"m04"`
	M05 float64 `json:"m05"`
	M06 float64 `json:"m06"`
	M07 float64 `json:"m07"`
	M08 float64 `json:"m08"`
	M09 float64 `json:"m09"`
	M10 float64 `json:"m10"`
	M11 float64 `json:"m11"`
	M12 float64 `json:"m12"`
	M13 float64 `json:"m13"`
	M14 float64 `json:"m14"`
	M15 float64 `json:"m15"`
	M16 float64 `json:"m16"`
	M17 float64 `json:"m17"`
	M18 float64 `json:"m18"`
	M19 float64 `json:"m19"`
	M20 float64 `json:"m20"`
	M21 float64 `json:"m21"`
	M22 float64 `json:"m22"`
	M23 float64 `json:"m23"`
	M24 float64 `json:"m24"`
	M25 float64 `json:"m25"`
	M26 float64 `json:"m26"`
	M27 float64 `json:"m27"`
	M28 float64 `json:"m28"`
	M29 float64 `json:"m29"`
	M30 float64 `json:"m30"`
	M31 float64 `json:"m31"`
	M32 float64 `json:"m32"`
	M33 float64 `json:"m33"`
	M34 float64 `json:"m34"`
	M35 float64 `json:"m35"`
	M36 float64 `json:"m36"`
	M37 float64 `json:"m37"`
	M38 float64 `json:"m38"`
	M39 float64 `json:"m39"`
	M40 float64 `json:"m40"`
	M41 float64 `json:"m41"`
	M42 float64 `json:"m42"`
	M43 float64 `json:"m43"`
	M44 float64 `json:"m44"`
	M45 float64 `json:"m45"`
	M46 float64 `json:"m46"`
	M47 float64 `json:"m47"`
	M48 float64 `json:"m48"`
	M49 float64 `json:"m49"`
	M50 float64 `json:"m50"`
	M51 float64 `json:"m51"`
	M52 float64 `json:"m52"`
	M53 float64 `json:"m53"`
	M54 float64 `json:"m54"`
	M55 float64 `json:"m55"`
	M56 float64 `json:"m56"`
	M57 float64 `json:"m57"`
	M58 float64 `json:"m58"`
	M59 float64 `json:"m59"`
}

func (l *LiveDataTime) setLiveDataTimeValue(name string, value float64) {
	switch name {
	case "m00":
		l.M00 += value
	case "m01":
		l.M01 += value
	case "m02":
		l.M02 += value
	case "m03":
		l.M03 += value
	case "m04":
		l.M04 += value
	case "m05":
		l.M05 += value
	case "m06":
		l.M06 += value
	case "m07":
		l.M07 += value
	case "m08":
		l.M08 += value
	case "m09":
		l.M09 += value
	case "m10":
		l.M10 += value
	case "m11":
		l.M11 += value
	case "m12":
		l.M12 += value
	case "m13":
		l.M13 += value
	case "m14":
		l.M14 += value
	case "m15":
		l.M15 += value
	case "m16":
		l.M16 += value
	case "m17":
		l.M17 += value
	case "m18":
		l.M18 += value
	case "m19":
		l.M19 += value
	case "m20":
		l.M20 += value
	case "m21":
		l.M21 += value
	case "m22":
		l.M22 += value
	case "m23":
		l.M23 += value
	case "m24":
		l.M24 += value
	case "m25":
		l.M25 += value
	case "m26":
		l.M26 += value
	case "m27":
		l.M27 += value
	case "m28":
		l.M28 += value
	case "m29":
		l.M29 += value
	case "m30":
		l.M30 += value
	case "m31":
		l.M30 += value
	case "m32":
		l.M32 += value
	case "m33":
		l.M33 += value
	case "m34":
		l.M34 += value
	case "m35":
		l.M35 += value
	case "m36":
		l.M36 += value
	case "m37":
		l.M37 += value
	case "m38":
		l.M38 += value
	case "m39":
		l.M39 += value
	case "m40":
		l.M40 += value
	case "m41":
		l.M41 += value
	case "m42":
		l.M42 += value
	case "m43":
		l.M43 += value
	case "m44":
		l.M44 += value
	case "m45":
		l.M45 += value
	case "46":
		l.M46 += value
	case "47":
		l.M47 += value
	case "m48":
		l.M48 += value
	case "m49":
		l.M49 += value
	case "m50":
		l.M50 += value
	case "51":
		l.M51 += value
	case "m52":
		l.M52 += value
	case "m53":
		l.M53 += value
	case "m54":
		l.M54 += value
	case "m55":
		l.M55 += value
	case "m56":
		l.M56 += value
	case "m57":
		l.M57 += value
	case "m58":
		l.M58 += value
	case "m59":
		l.M59 += value
	}
}

func main()  {
	dataTime := new(LiveDataTime)
	dataTime.setLiveDataTimeValue("m59", 1)
	dataTime.setLiveDataTimeValue("m58", 2)
	arr := []int{1,2,3,4,5,6,7,8}
	var arrStr []string
	length := len(arr)
	var cond string
	var i int
	for _, a := range arr {
		if i == 0 {
			i++
			cond += fmt.Sprintf("datapoint_id = %d", a)
		} else {
			i++
			cond += fmt.Sprintf(" or datapoint_id = %d", a)
		}
		if length > 10 && i == 10 {
			cond += ";"
			i = 0
		}
	}
	arrStr = strings.Split(cond, ";")
	fmt.Println(arrStr[0])
}