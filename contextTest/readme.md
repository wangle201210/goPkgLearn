## pkg path
https://github.com/golang/go/tree/master/src/context

## learn from
[简书中文介绍](https://www.cnblogs.com/kaichenkai/p/11352377.html)  
参考下别人的写法

## 实际操作过程

1. context.WithValue  
    用在上下文传值上，比如方法之间的参数传递（实际情况肯定不会大么简单）

1. context.WithCancel  
    用在控制自己想精准退出的位置
    
1. context.WithTimeout  
    用在判定超时上，比如网络请求等  
    本意是过*秒后结束
    
1. context.WithDeadline  
    和WithTimeout类似，不过第二个参数传入的是过期时间  
    即到**时结束
    
