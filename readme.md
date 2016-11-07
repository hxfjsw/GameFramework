#### 配置方法
修改 conf/game.ini 文件
```
[net]
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
- 当收到一个连接
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
timer_after(ms,func)
```
- 设置一个间隔时间定时器
```
timer_tick(ms,func)
```
