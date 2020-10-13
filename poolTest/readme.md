## pkg path
自己实现一个链接池

## learn from
和[wbofeng](https://github.com/flutterWang) 一起探讨学习  

## 实现方法
本质上和sync/pool 类似
只是实现逻辑简单些
```go
// 首先去new一个pool
p, err := pool.New(size, creatorFunc, interval)
// 获取一个池内元素
element, err := p.Get()
// 扔回链接池
p.Put(element)
// 关闭整个链接池
p.close()
```
