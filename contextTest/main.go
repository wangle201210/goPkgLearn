package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

// 在超时上使用
func usedTimeout(i int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(i))
	defer cancel()
	fmt.Println("ctx初始化了，看下输入时间是否大于3s")
	select {
	case <-ctx.Done():
		fmt.Println("超时了，被允许的时间小于了3s")
	case <-time.After(time.Second * 3):
		fmt.Println("3 秒过去了，没有超时！")
	}
	return
}

// 传递参数时使用
func usedValue(org int) {
	ctx := context.WithValue(context.Background(), "data", org)
	v := addValue(ctx, 1)
	fmt.Printf("%d+1=%d",org,v)
}

func addValue(ctx context.Context, add int) int {
	org := ctx.Value("data").(int)
	return org + add
}

// 想在某个确定的地方取消时使用
func usedCancel(stop int) {
	if stop > 100 {
		fmt.Println("请勿大于100")
		return
	}
	c := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for i := 0; i < 100; i++ {
			if i == stop {
				cancel()
			}
			c <- i
			fmt.Printf("input:  %d \n",i)
		}
	}()
	for {
		select {
		case res := <-c:
			fmt.Printf("output: %d \n", res)
		case <-ctx.Done():
			return
		}
	}
}

func usedDeadline(i int) {
	addTime := time.Second*time.Duration(i)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(addTime))
	defer cancel()
	fmt.Println("ctx初始化了")
	index := 0
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			index++
			fmt.Printf("%ds has passed\n",index)
		case <-ctx.Done():
			fmt.Println("ctx 到了取消时间了")
			return
		}
	}

}
func main() {
	usedDeadline(10)
	// 打印到小于参数时停止
	//usedCancel(200)
	// 将传入的参数加一
	//usedValue(99)
	// 传入参数小于3则会超时
	//usedTimeout(2)
}
