package main

import (
	"fmt"
	"time"
)

func main() {
	//// print 4 3 2 1 0
	//for i := 0; i < 5; i++ {
	//	defer fmt.Println(i)
	//}
	//// print 5 5 5 5 5
	//for i := 0; i < 5; i++ {
	//	defer func(i *int) {
	//		fmt.Println(*i)
	//	}(&i)
	//}
	//// 0s
	//startAt := time.Now()
	//defer fmt.Println(time.Since(startAt))
	//time.Sleep(time.Second)
	// 1.***s
	startAt1 := time.Now()
	defer func() {
		fmt.Println(time.Since(startAt1))
	}()
	time.Sleep(time.Second)

}