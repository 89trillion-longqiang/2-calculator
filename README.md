#计算器

#整体框架
```
使用Gin框架完成一个接口：请求参数一个字符串，包含正整数、加(+)、减(-)、
乘(*)、除(/)的算数表达式(括号除外)，返回表达式的计算结果。表达式仅包含
非负整数， +， - ，*，/ 四种运算符和空格 。 整数除法仅保留整数部分。
```

#目录结构
```
├── PreTest                             #压力测试
│   ├── calculator.py
│   └── report_calculator.html
├── README.md                           #介绍
├── calculator_test.go                  #单元测试
├── go.mod
├── handle                              #处理router的请求
│   ├── handle.go
│   └── routerCtrl.go
├── main.go                             #代码入口
├── router                              #路由转发
│   └── router.go
└── service                             #处理计算器逻辑
    └── cal.go
```


#代码逻辑分层  calculator
| 层     | 文件夹|主要职责 |调用关系|
| :----: | :----|:---- | :-----|
|router  | /router|路由转发 |调用handle|
|handle   | /handle|处理路由 |被router调用 调用service|
|service | /service|处理通用逻辑| 被handle调用 |


#接口设计
```
接口地址 
/C1/calculator 
```
### 请求方式
GET
### 请求示例
```
http://127.0.0.1:9000/C1/calculator?exp=3+3*4/2
```

### 参数  说明

``` 
exp 类型 string 输入需要一个字符串包含正整数、加(+)、减(-)、乘(*) 、除(/)的算数表达式(括号除外) 
```

```
成功示例 
{
  "condition":"pass",
  "result":9
} 
错误示列 
{
  "condition":"error",
  "result":"exp error"
}
```

#第三方库
```
github.com/gin-gonic/gin  开发框架
```