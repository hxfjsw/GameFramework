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

- 当收到一个连接

- 当收到网络消息

- 当一个连接关闭