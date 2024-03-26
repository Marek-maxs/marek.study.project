package main

import (
	"context"
	"fmt"
"time"

cache "github.com/chenmingyong0423/go-generics-cache"
)

func main() {
	data := cache.NewSimpleCache[int, string](0, 3*time.Second)
	err := data.Set(context.Background(), 1, "test cache info", cache.WithExpiration(time.Millisecond))
	fmt.Println(err)
	res, err := data.Get(context.Background(), 1)
	fmt.Println(err)
	fmt.Println(res)
}