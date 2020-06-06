class Websocket {

    /**
     * 连接句柄
     */
    ws = null;

    /**
     * 最大重连次数
     */
    MAX_RECONNECT = 5;

    /**
     * 当前重连次数
     */
    CURRENT_RECONNECT_COUNT = 0;

    /**
     * 心跳定时器
     */
    timer = null;

    /**
     * 连接状态
     */
    status = false;

    constructor(addr, onmessage) {
        this.openConnect(addr, onmessage)
    }

    openConnect(addr, onmessage) {
        let ws = new WebSocket(addr);
        this.ws = ws;
        let that = this;
        ws.onopen = function() {
            that.CURRENT_RECONNECT_COUNT = 0;
            that.status = true
            that.timer = setInterval(function () {
                let data = {
                    type: 3,
                    data: "PING"
                };
                ws.send(JSON.stringify(data));
            }, 10000);
        };
        ws.onmessage = onmessage
        ws.onclose = function() {
            that.clearHeartbeatTimer()
            that.status = false
            setTimeout(function () {
                if (that.CURRENT_RECONNECT_COUNT++ < that.MAX_RECONNECT) {
                    that.openConnect(addr, onmessage);
                }
            }, 3000)
        };
    }

    clearHeartbeatTimer() {
        if (this.timer != null) {
            clearInterval(this.timer);
        }
    }

    pushToServer(content) {
        this.ws.send(content)
    }
}
