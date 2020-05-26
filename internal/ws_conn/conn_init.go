package ws_conn

import (
	"github.com/gorilla/websocket"
	"go-IM/internal/logic/service"
	"log"
	"net/http"
	"strconv"
)

func StartServer(addr string) {
	http.HandleFunc("/ws", wsHandler)
	log.Print("websocket server start run in ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Panic(err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 协议升级，Hold连接
func wsHandler(w http.ResponseWriter, r *http.Request) {
	appId, _ := strconv.ParseInt(r.Header.Get("app_id"), 10, 64)
	userId, _ := strconv.ParseInt(r.Header.Get("user_id"), 10, 64)
	deviceId, _ := strconv.ParseInt(r.Header.Get("device_id"), 10, 64)
	token := r.Header.Get("token")
	requestId, _ := strconv.Atoi(r.Header.Get("request_id"))

	if appId == 0 || userId == 0 || deviceId == 0 || token == "" || requestId == 0 {
		_, _ = w.Write([]byte("error unauthorized"))
	}

	err := service.AuthService.SignIn(appId, userId, deviceId, token, "", 0)
	if err != nil {
		log.Print(err)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	conn, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Print(e)
		return
	}

	// 断开这个设备之前的连接
	preCtx := load(deviceId)
	if preCtx != nil {
		preCtx.DeviceId = PreConn
	}

	ctx := NewConnContext(conn, appId, deviceId, userId)
	store(deviceId, ctx)
	ctx.DoConn()
}
