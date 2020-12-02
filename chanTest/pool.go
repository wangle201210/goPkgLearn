package main

import (
	"errors"
	"time"

	"sync/atomic"
)

var (
	errSizeTooSmall = errors.New("[error] init: size of the pool is too small")
	errPoolClosed   = errors.New("[error] get: pool is closed")
	errTimeout      = errors.New("[error] get: timeout")
	errAllExpire    = errors.New("[error] pool: all expire")
)

// Pool -
type Pool struct {
	active   int64
	max      int64
	res      chan Element
	signal   chan bool
	close    chan bool
	interval time.Duration
	creator  func() Element
}

// Element -
type Element interface {
	Expire() bool
}

// New -
func New(max int64, timer time.Duration, fn func() Element) (*Pool, error) {
	if max <= 0 {
		return nil, errSizeTooSmall
	}

	pool := &Pool{
		active:   0,
		max:      max,
		res:      make(chan Element, max),
		close:    make(chan bool),
		signal:   make(chan bool),
		interval: timer,
		creator:  fn,
	}

	for i := 0; i < int(pool.max/2); i++ {
		pool.res <- pool.creator()
		pool.active++
	}

	go pool.start()

	return pool, nil
}

func (p *Pool) start() {
	for {
		select {
		case <-p.close:
			for {
				if int64(len(p.res)) == p.active {
					close(p.signal)
					close(p.res)
					return
				}
			}
		case <-p.signal:
			if p.active < p.max {
				p.res <- p.creator()
				p.active++
			}
		}
	}
}

func (p *Pool) check(element Element) (Element, error) {
	if element.Expire() {
		atomic.AddInt64(&p.active, -1)
		if p.active == 0 {
			return nil, errAllExpire
		}

		ele, err := p.Get()
		if err != nil {
			return nil, errAllExpire
		}

		return ele, nil
	}

	return element, nil
}

// Get -
func (p *Pool) Get() (Element, error) {
	select {
	case element := <-p.res:
		ele, err := p.check(element)
		if err != nil {
			return nil, err
		}
		return ele, nil
	case <-p.close:
		return nil, errPoolClosed
	case p.signal <- true:
	}

	ticker := time.NewTimer(p.interval)
	select {
	case element := <-p.res:
		ticker.Stop()
		return element, nil
	case <-ticker.C:
		return nil, errTimeout
	case <-p.close:
		return nil, errPoolClosed
	}
}

// Put -
func (p *Pool) Put(element Element) {
	p.res <- element
}

// Close -
func (p *Pool) Close() {
	close(p.close)
}
