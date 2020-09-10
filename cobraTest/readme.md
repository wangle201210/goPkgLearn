## pkg path
https://github.com/spf13/cobra

## learn from
[简书](https://www.cnblogs.com/sparkdev/p/10856077.html)

## 实际操作过程
1. 安装cobra  
    `go get -u github.com/spf13/cobra/cobra`
    
1. 创建项目(这里简书上那篇文章讲的不明确)  
    命令行的格式  
    `cobra init 包名 --pkg-name=包地址`  
    我试过，如果不加 --pkg-name 就一直出问题～  
    我的完整命令如下  
    `cobra init cobraTest --pkg-name=github.com/wangle201210/goPkgLearn/cobra`

1. 添加命令
    `cobra add say`  
    在init方法中添加
    `sayCmd.Flags().StringVarP(&say,"str","s","hello word","enter wath you want to say")`  
    上面的意思是有变量say去接受输入的--str或者-s后面的内容，如果没有跟内容就默认取"hello world"
    -s是--str的简写  
1. 剩下的就看 [简书](https://www.cnblogs.com/sparkdev/p/10856077.html) 里面的介绍了
