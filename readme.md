# 异步、并行、高性能、可热更新的网络通信引擎

## 架构
- 利用Golang超强的并发网络处理能力去处理网络层的业务,用js处理业务逻辑以实现热更新
- 热更新的方法就是向进程发送 signal-usr2
- 基于goroutine,可以轻松实现异步

## 对比
- 与boost.asio + lua 相比服务器更稳定,极少出现coredump 服务器开发效率也远比c/c++高
- 与node.js 相比拥有更高的性能,基于goroutine方式实现的异步在使用中也远比基于函数闭包的异步方便,不会出现`回调地狱`

## 安装方法
```
yum install go
go get github.com/robertkrimen/otto
go get github.com/Unknwon/goconfig
go get github.com/op/go-logging
bash build.sh
```

#### 配置方法
修改 conf/game.ini 文件
```
[log]
file = ./logs/common.log
level = DEBUG

[javascript]
file = ./scripts/main.js

[http]
ip = 0.0.0.0
port = 80

[tcp]
ip = 0.0.0.0
port = 2001

[websocket]
ip = 0.0.0.0
port = 2002

[mysql]
ip = 127.0.0.1
port = 3306
user = root
pwd = 123456
db = test

[redis]
ip = 127.0.0.1
port = 6379
```

#### 事件回调函数

- 当进程启动时
```
funtion onStart()
```
- 当进程关闭
```
function onShutdown()
```
- 当收到一个网络连接
```
function onConnect(fd)
```
- 当收到网络消息
```
function onMessage(fd,msg)
```
- 当一个连接关闭
```
function onClose(fd)
```
- 收到一个HTTP请求
```
function onRequest(request) (result string)
返回一个string类型 直接发回给浏览器
```

#### 主动调用函数

- 指定时间后调用函数
```
function timer_after(ms,func)
```
- 设置一个间隔时间定时器
```
function timer_tick(ms,func)
```
- 向某个网络连接发送消息
```
function sendToFd(fd,msg)
```
- 全服广播一个消息
```
function boardcast(msg)
```
- 主动关闭服务器
```
function exit()
```
- 打印一条日志
```
function log(msg)
```
