package ws_handler

import (
	"github.com/gorilla/websocket"
	"go-IM/logic/service"
	"go-IM/pkg/defs"
	"log"
	"net/http"
)

var pushChan *chan defs.MessageItem

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 协议升级，Hold连接
func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Print(e)
		return
	}
	ctx := NewConnContext(conn)
	ctx.DoConn()

	messageChan := make(chan defs.MessageItem)
	pushChan = &messageChan
	service.MessageService.PushChan = pushChan
	DoSend()
}

// 内部通信
func DoSend() {
	go func() {
		for {
			message := <-*pushChan
			ctx := load(message.ReceiverDeviceId)
			if ctx != nil {
				ctx.Output(defs.PackageType_RT_USER, 0, nil, message)
			}
		}
	}()
}
