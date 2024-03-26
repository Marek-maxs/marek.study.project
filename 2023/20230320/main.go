package main

import (
	"fmt"
)

func main()  {
	//i := 0
	//for j := 0; j < 10; j++ {
	//	for i = 0; i < 10; i++ {
	//		if model.CheckMeterExists(fmt.Sprintf("%d", i)) {
	//			continue
	//		}
	//		fmt.Println("start:", i)
	//		key := fmt.Sprintf("%d",i)
	//		meter := map[string]string{key:fmt.Sprintf("%d",i)}
	//		model.SetMetersMap(meter)
	//		go func(n int) {
	//			t := time.After(1 * time.Second)
	//			_, _ = <- t
	//			model.ClearMetersMapByKey(fmt.Sprintf("%d", n))
	//			fmt.Println("done:", n)
	//		}(i)
	//	}
	//	i = 0
	//
	//}
	//data := make(map[string][]string)
	//pov := make([]string, 0, 1)
	//pov = []string{"11", "22"}
	//data["1"] = pov
	//pov = []string{"33", "44"}
	//data["1"] = pov
	//fmt.Println(data)
	//wattHour, err := strconv.ParseFloat("-0.0001", 64)
	//fmt.Println(err)
	//fmt.Println(wattHour)
	pool := make(chan string, 200)

	for i := 0; i < 10; i++ {
		pool <- fmt.Sprintf("%d", i)
	}
	for i := 0; i < 10; i++ {
		pool <- fmt.Sprintf("%d", i)
	}
	n := 0
	for {
		n++
		if n == 20 {
			break
		}
		msg := <- pool
		fmt.Println(msg)
	}
	<- pool/1``
	fmt.Println("-----33-----")
	fmt.Println(len(pool))
}