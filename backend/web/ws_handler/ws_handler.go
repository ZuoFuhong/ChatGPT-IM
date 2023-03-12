package ws_handler

import (
	"ChatGPT-IM/backend/logic/service"
	"ChatGPT-IM/backend/pkg/defs"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func init() {
	mc := make(chan defs.MessageItem, 100)
	service.MessageService.PushChan = mc
	DoConsume(mc)
}

// WSHandler 协议升级，Hold连接
func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Print(e)
		return
	}
	ctx := NewConnContext(conn)
	ctx.DoConn()
}

// DoConsume 内部通信
func DoConsume(mc chan defs.MessageItem) {
	go func() {
		for {
			message := <-mc
			ctx := load(message.ReceiverDeviceId)
			if ctx != nil {
				ctx.Output(defs.PackagetypeRtUser, 0, nil, message)
			}
		}
	}()
}
