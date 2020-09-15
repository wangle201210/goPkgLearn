## pkg path
https://github.com/golang/go/tree/master/src/sync

## learn from
[参考文档](https://www.cnblogs.com/single-dont/p/13401306.html)

## 实际操作过程
1. mutex  
    共两个方法
    ```go
       func (m *Mutex) Lock() // 锁
       func (m *Mutex) Unlock() // 解锁
    ```
    切记锁的竞争是公平(随机)的，任何一个都可能先获取到锁，我在测试用例中有测试  
    不仅是sync.Mutex,sync.RWMutex也是一样的
    ```bash
    我在第4个chan 处被锁了
    过去了3s，我在4处解锁了
    我在第1个chan 处被锁了
    过去了3s，我在1处解锁了
    我在第3个chan 处被锁了
    过去了3s，我在3处解锁了
    我在第2个chan 处被锁了
    过去了3s，我在2处解锁了
    ```
1. RWMutex  
    共四个方法
    ```go
    func (rw *RWMutex) Lock() // 写锁
    func (rw *RWMutex) Unlock() // 写解锁
    
    func (rw *RWMutex) RLock() // 读锁
    func (rw *RWMutex) RUnlock() // 读解锁
    ```
    着重理解下  
    可以多个一起读  
    写锁相互排斥  
    有读不可写  
    有写不可读
    
    我使用了两个方法验证加读写锁后的安全  
    其中`fun rwMutex(i int)` 使用了RWMutex锁，开协程并发的累加到10000，同时开协程累减到10000，最终得到的值为0，证明加锁后是独写安全的   
    `fun rwWithoutMutex(i int)` 没有使用了RWMutex锁，同上一方法先同时累加累减，得到结果不为预期值0，所以独写不安全
 
1. WaitGroup
    共三个方法
    ```go
    func (wg *WaitGroup) Add(i int) // 往待处理组里面加i
    func (wg *WaitGroup) Done() // 待处理组里的待处理数-1
    func (wg *WaitGroup) Wait() // 等待待处理组里面的数字减小到0，证明就已经处理完了
    ```
    在我看来这个主要的作用就是‘计数’待处理的事项，直到处理好了，再开始做接下来的事情