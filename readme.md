## 架构
利用Golang超强的并发网络处理能力去处理网络层的业务,用js处理业务逻辑以实现热更新
热更新的方法就是向进程发送 signal-usr2

## 安装方法
yum install go
go get 

#### 配置方法
修改 conf/game.ini 文件
```
[tcp]
ip = 0.0.0.0
port = 2001

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
function onClose(fd)

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
