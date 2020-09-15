package main

import (
	"fmt"
	"sync"
	"time"
)

type info struct {
	sync.RWMutex
	data int
}
func mutex(conut int) {
	c := make(chan struct{},conut)
	lock := sync.Mutex{}
	// 锁的竞争是公平（随机）的，任何一个都可能先获取到锁
	// 但是goroutine不是公平的。。。
	for i := 1; i < conut; i++ {
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Printf("我在第%d个chan 处被锁了\n",i)
			time.Sleep(time.Second )
			fmt.Printf("过去了1s，我在%d处解锁了\n",i)
			c <- struct {
			}{}
		}(i)
	}
	for i := 1; i < conut; i++ {
		<-c
	}
}

func rwMutex(count int)  {
	c := make(chan struct{},count * 3)
	l := info{data:0}
	for i := 0; i < count; i++ {
		go func() {
			l.RLock()
			d := l.data
			fmt.Printf("我读取到了data，值为:%d\n",d)
			l.RUnlock()
			c <- struct{}{}
		}()
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			l.Lock()
			l.data += i
			fmt.Printf("我把data的值加了%d变成了%d\n",i,l.data)
			l.Unlock()
			c <- struct{}{}
		}(i)
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			l.Lock()
			l.data -= i
			fmt.Printf("我把data的值减了%d变成了%d\n",i,l.data)
			l.Unlock()
			c <- struct{}{}
		}(i)
	}
	for i := 0; i < count * 3; i++ {
		<-c
	}
	fmt.Printf("data的最终结果应该为0，实际结果为：%d",l.data)
}

func rwWithoutMutex(count int)  {
	c := make(chan struct{},count * 3)
	l := 0
	for i := 0; i < count; i++ {
		go func() {
			fmt.Printf("我读取到了data，值为:%d\n",l)
			c <- struct{}{}
		}()
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			l += i
			fmt.Printf("我把data的值加了%d变成了%d\n",i,l)
			c <- struct{}{}
		}(i)
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			l -= i
			fmt.Printf("我把data的值减了%d变成了%d\n",i,l)
			c <- struct{}{}
		}(i)
	}
	for i := 0; i < count * 3; i++ {
		<-c
	}
	fmt.Printf("不安全读写时data的最终结果应该为0，实际结果为：%d",l)
}

func wg() {
	w := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		w.Add(1)
		fmt.Println("我往等待处理的组里加了1")
		go func() {
			fmt.Println("处理好了，往待处理的组里减少1")
			defer w.Done()
			time.Sleep(time.Second)
		}()
	}
	w.Wait()
	fmt.Println("我等待待处理组里面把东西处理完，然后退出程序")
}
func main() {
	// 得到的结果为0，证明加锁安全
	//rwMutex(10000)
	// 得到结果不为0，证明不加锁不安全
	rwWithoutMutex(10000)
	//mutex(1000)
	//wg()
}

