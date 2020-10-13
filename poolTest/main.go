package main

import (
	"errors"
	"sync/atomic"
	"time"
)

var (
	errEleTimeOut    = errors.New("该链接已超时")
	errCreateTimeOut = errors.New("创建元素超时")
	errPoolSize      = errors.New("池容量必须大于0")
)

type Pool struct {
	size     int64 //可以存放的元素个数
	count    int64 // 池里面已经创建的个数
	res      chan Element
	close    chan bool      //关闭池时使用
	signal   chan bool      //需要往res里面创建新元素时使用
	creator  func() Element //自定义创建res元素的方法
	interval time.Duration  // 创建元素前的等待时间，超过则报超时
}

type Element interface {
	Expire() bool //实现这个方法就能当成pool的元素
}

func New(size int64, creator func() Element, interval time.Duration) (*Pool, error) {
	if size <= 0 {
		return nil, errPoolSize
	}
	var p = &Pool{
		size:     size,
		count:    0,
		res:      make(chan Element, size),
		close:    make(chan bool),
		signal:   make(chan bool),
		creator:  creator,
		interval: interval,
	}
	// 先往池里面填充一半的内容，以备之后使用
	for i := 0; i < int(size)/2; i++ {
		p.res <- p.creator()
		p.count++
	}
	go p.start()
	return p, nil
}


// 判断下当前element能不能使用,不能使用则扔掉
func (p *Pool) check(e Element) (Element, error) {
	if e.Expire() {
		atomic.AddInt64(&p.count,-1)
		element, err := p.Get()
		if err != nil {
			return nil, err
		}
		return element, nil
	}
	return e,nil
}

func (p *Pool) start() {
	select {
	case <-p.close: //关闭时，先等待全部归还
		for {
			// 当前res元素数量==创建了的总数时，才能正式关闭
			if int(p.count) == len(p.res) {
				close(p.res)
				close(p.signal)
				close(p.close)
			}
		}
	case <-p.signal: //创建新元素
		if p.count < p.size {
			p.res <- p.creator()
			p.count++
		}
	}
}

// 获取池里的一个元素
func (p *Pool) Get() (Element, error) {
	// 如果res里面有元素，就取出来用
	// 否则传入信号量，让pool新建一个元素
	select {
	case element := <-p.res:
		ele, err := p.check(element)
		if err != nil {
			return nil, err
		}
		return ele, nil
	case p.signal <- true:
	}
	select {
	case r := <-p.res:
		return r, nil
	case <-time.After(p.interval): //创建超时了
		return nil, errCreateTimeOut
	}
}

// 放回元素
func (p *Pool) Put(ele Element) {
	p.res <- ele
}

func (p *Pool) Close() {
	p.close <- true
}

