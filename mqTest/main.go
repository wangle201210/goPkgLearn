package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)
// 每个消息队列里面可以存放的消息数量
const size = 1000

// 订阅用户
type user struct {
	name string
	//...
}

// 消息
type message struct {
	msg string
	//...
}

// 具体消息的队列
type queue struct {
	m chan message
}

// 容器
type broker struct {
	sync.RWMutex //保证topic的读写安全

	topics map[string]queue //用来放不同主题的消息
}

// 订阅者
type subscriber struct {
	sync.RWMutex

	users map[string][]user // 用户列表 topic => userList
}

//var subscriber chan user

func addSub(name, topic string) {
	u := user{name: name}
	subs.Lock()
	subs.users[topic] = append(subs.users[topic], u)
	subs.Unlock()
}

func addMsg(msg string, topic string) {
	fmt.Printf("向 %s topic中添加了 %s 消息\n", topic, msg)
	if brok.topics[topic].m == nil {
		q := make(chan message,size)
		t := queue{m: q}
		brok.Lock()
		brok.topics[topic] = t
		brok.Unlock()
	}
	m := message{msg: msg}
	brok.topics[topic].m <- m
}

var subs = subscriber{
	users: make(map[string][]user),
}

var brok = broker{
	topics: make(map[string]queue),
}


func doAddUser() {
	// 添加订阅者，假设有两类不同的订阅者
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("第%d位用户", i)
		addSub(name, "t0")
	}

	for i := 10; i < 20; i++ {
		name := fmt.Sprintf("第%d位用户", i)
		addSub(name, "t1")
	}
}

func doAddMsg(n int) {
	for i := 0; i < n; i++ {
		tid := i%2
		t := fmt.Sprintf("t%d", tid)
		addMsg(strconv.Itoa(i), t)
		// 每秒发一次消息过去
		time.Sleep(time.Second)
	}
}

func custom() {
	// 遍历每一个topic,并且发送给订阅了的用户
	for k, v := range brok.topics {
		go func() {
			select {
			case m := <-v.m:
				for _, vu := range subs.users["t1"] {
					go send(vu, k, m)
				}
			default:
				time.Sleep(time.Second)
			}
		}()
	}
}

func send(u user,t string, m message) {
	fmt.Printf("给 %s 用户在%s主题中发送了 %s 消息\n", u.name,t, m.msg)
}

func main() {
	// 添加消费者
	go doAddUser()
	// 每隔1s发送一次消息，发送n次
	go doAddMsg(10)
	// 每隔2s去消费一次
	for  {
		custom()
		time.Sleep(time.Second * 2)
	}
}
