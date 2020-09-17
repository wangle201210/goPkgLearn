## pkg path
https://github.com/tealeg/xlsx

## learn from
[参考文档](https://www.jianshu.com/p/c1753d517fa0)

## 原因
一个同学参加数学建模，需要处理excel数据

## 需求
将 "操作变量"工作簿 中的每列的数据中的0值，替换为该列的平均值  
其中4-43行为一份数据45-84行为第二份数据  

## 处理
就简单的将数据读取到slice里面然后  
替换掉对应位置的数据  
然后重新生成一个新文件