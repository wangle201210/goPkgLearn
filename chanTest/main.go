package main

import (
	"fmt"
	"time"
)

type m struct {
	c chan int
	data string
}
func main() {
	// 正常情况close需要放在写入ch的goroutine里面去
	// 当close不放在写入ch的goroutine时会产生以下问题
	// 如果缓冲区容量小于10-first就会报错说'send on closed channel'
	// 假如ch的缓冲区能放下没有被接收的数据，就不会报错
	// 但是这个报错有时候不会被触发（反复go run . 多次后就可能有一次没报错）

	// 假设容量为2 则报错（有大概80%机会panic）
	size := 2
	// 0-4号元素被直接被消费掉了
	// 5-6（4+2-1）号元素被放到缓冲区
	// 7-9被放到recvq阻塞队列中去了
	// 此时关闭ch
	// 然后再去读取ch
	// 有时候（大多数情况）会报错，有时候不会报错

	// 假设容量为8 则不会报错
	//size := 8
	// 因为缓冲区有足够的地方放数据
	// 不会有元素挂载到recvq里面
	ch := make(chan int,size)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Printf("往 ch 中输入了%d\n",i)

		}
		//close(ch)
	}()
	first := 4
	go func() {
		for i := 0; i < first; i++ {
			x,ok := <-ch
			fmt.Printf("%t,ch 中输出了%d\n",ok,x)
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("=====================")
	close(ch)
	for i := 0; i < 10 - first; i++ {
		x,ok := <-ch
		fmt.Printf("%t,ch 中输出了%d\n",ok,x)
	}
}


