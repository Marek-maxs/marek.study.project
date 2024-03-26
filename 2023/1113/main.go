package main

import (
	"context"
	"fmt"
	"time"
)

/**
*
* Author: Marek
* Date: 2023-11-13 9:59
* Email: 364021318@qq.com
*
 */

// 原地数组去重
// 非排序数组
func removeDuplication_map(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

func main() {
	a := true

	if !a {
		fmt.Println("lock see!")
	}
	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	fmt.Println("start at ", time.Now().Format("2006-01-02 15:04:05"))
	go goroutinest(ctx)

	time.Sleep(10 * time.Second) // 主程序用sleep阻塞住
	fmt.Println("stop at ", time.Now().Format("2006-01-02 15:04:05"))
}

func goroutinest(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():  // 因为ctx带超时参数，当时间期限到了之后就会走到这里退出协程
			fmt.Println("Done. ", time.Now().Format("2006-01-02 15:04:05"))
			return
		default:            // 协程循环执行for，当ctx.Done()无信号时总是走到Default分支
			// fmt.Println("case default ", time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(time.Second)
		}
	}
}