/**
 * Created by huangxiufeng on 16/11/6.
 */

/**
 * 当进程启动
 */
function onStart() {
    log("starting.....");
    timer_after(3000, "log('after clock!!');");
    timer_tick(1000, "onTick();");


}


function onTick() {
    log("tick clock!!")
}

/**
 * 当进程关闭
 */
function onShutdown() {
    log("onShutdown");
    exit();
}

/**
 * 收到消息
 * @param fd
 * @param message
 */
function onMessage(fd, message) {
    var ip = fdToIp(fd);
    log("收到消息 fd=" + fd + " message=" + message);
    sendToFd(fd, "recv:" + message);
    broadcast("msg from" + fd + ":" + msg);
}

/**
 * 连接打开
 * @param fd
 */
function onConnect(fd) {
    sendToFd(fd, "hello!");
}


function onClose(fd) {
    log("连接断开 fd=" + fd);
}


