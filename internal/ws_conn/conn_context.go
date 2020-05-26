package ws_conn

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go-IM/internal/logic/service"
	"go-IM/pkg/defs"
	"go-IM/pkg/util"
	"log"
	"time"
)

const PreConn = -1 // 设备第二次重连时，标记设备的上一条连接

type ConnContext struct {
	Conn     *websocket.Conn
	AppId    int64
	DeviceId int64
	UserId   int64
}

func NewConnContext(conn *websocket.Conn, appId, deviceId, userId int64) *ConnContext {
	return &ConnContext{
		Conn:     conn,
		AppId:    appId,
		DeviceId: deviceId,
		UserId:   userId,
	}
}

func (ctx *ConnContext) DoConn() {
	util.RecoverPanic()

	for {
		err := ctx.Conn.SetReadDeadline(time.Now().Add(12 * time.Minute))
		if err != nil {
			fmt.Print(err)
			return
		}
		_, data, err := ctx.Conn.ReadMessage()
		if err != nil {
			fmt.Print(err)
			return
		}
		ctx.HandlePackage(data)
	}
}

func (ctx *ConnContext) HandlePackage(data []byte) {
	var input defs.Input
	err := json.Unmarshal(data, &input)
	if err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	switch input.Type {
	case defs.PackageType_SYNC:
		ctx.Sync(input)
	case defs.PackageType_HEARTBEAT:
		ctx.Heartbeat(input)
	case defs.PackageType_MESSAGE:
		ctx.MessageACK(input)
	}
}

// 离线消息同步
func (ctx *ConnContext) Sync(input defs.Input) {
	var sync defs.SyncInput
	err := json.Unmarshal(input.Data, &sync)
	if err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	messages, err := service.MessageService.ListByUserIdAndSeq(ctx.AppId, ctx.UserId, sync.Seq)
	var syncOutput defs.SyncOutput
	if err == nil {
		syncOutput = defs.SyncOutput{Messages: *messages}
	}

	ctx.Output(defs.PackageType_SYNC, input.RequestId, err, &syncOutput)
}

func (ctx *ConnContext) Heartbeat(input defs.Input) {
	ctx.Output(defs.PackageType_HEARTBEAT, input.RequestId, nil, nil)
	log.Print("heartbeat ", " device_id ", ctx.DeviceId, " user_id ", ctx.UserId)
}

// 消息回执
func (ctx *ConnContext) MessageACK(input defs.Input) {
	var messageACK defs.MessageACK
	err := json.Unmarshal(input.Data, &messageACK)
	if err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	err = service.DeviceAckService.Update(ctx.DeviceId, messageACK.DeviceAck)
	if err != nil {
		log.Print(err)
	}
}

func (ctx *ConnContext) Output(pt defs.PackageType, requestId int, err error, message *defs.SyncOutput) {
	var output = defs.Output{
		Type:      pt,
		RequestId: requestId,
	}
	if err != nil {
		output.Code = 1
		output.Message = err.Error()
	}
	if message != nil {
		msgBytes, err := json.Marshal(message)
		if err != nil {
			log.Print(err)
			return
		}
		output.Data = msgBytes
	}
	outputBytes, err := json.Marshal(&output)
	if err != nil {
		log.Print(err)
		return
	}
	err = ctx.Conn.WriteMessage(websocket.TextMessage, outputBytes)
	if err != nil {
		log.Print(err)
		return
	}
}

// 释放TCP连接
func (ctx *ConnContext) Release() {
	e := ctx.Conn.Close()
	if e != nil {
		log.Print(e)
	}
	// 设备下线
	if ctx.DeviceId != PreConn {
		delete(ctx.DeviceId)
		_ = service.DeviceService.Offline(ctx.AppId, ctx.UserId, ctx.DeviceId)
	}
}
