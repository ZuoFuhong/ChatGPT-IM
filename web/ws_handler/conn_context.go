package ws_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go-IM/logic/service"
	"go-IM/pkg/defs"
	"go-IM/pkg/util"
	"log"
	"strconv"
	"time"
)

const PreConn = -1 // 设备第二次重连时，标记设备的上一条连接

type ConnContext struct {
	Conn     *websocket.Conn
	AppId    int64
	DeviceId int64
	UserId   int64
}

func NewConnContext(conn *websocket.Conn) *ConnContext {
	return &ConnContext{
		Conn: conn,
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
	case defs.PackageType_SignIn:
		ctx.SignIn(input)
	case defs.PackageType_SYNC:
		ctx.Sync(input)
	case defs.PackageType_HEARTBEAT:
		ctx.Heartbeat(input)
	case defs.PackageType_MESSAGE_ACK:
		ctx.MessageACK(input)
	case defs.PackageType_RT_USER:
		ctx.SendToUser(input)
	}
}

// 登录鉴权
func (ctx *ConnContext) SignIn(input defs.Input) {
	signIn := new(defs.SignIn)
	err := json.Unmarshal([]byte(input.Data), &signIn)
	if err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	appId, _ := strconv.ParseInt(signIn.AppId, 10, 64)
	userId, _ := strconv.ParseInt(signIn.UserId, 10, 64)
	deviceId, _ := strconv.ParseInt(signIn.DeviceId, 10, 64)
	err = service.AuthService.SignIn(appId, userId, deviceId, signIn.Token, "", 0)
	if err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	ctx.AppId = appId
	ctx.UserId = userId
	ctx.DeviceId = deviceId

	// 断开这个设备之前的连接
	preCtx := load(ctx.DeviceId)
	if preCtx != nil {
		preCtx.DeviceId = PreConn
	}
	store(ctx.DeviceId, ctx)
	ctx.Output(defs.PackageType_SignIn, input.RequestId, err, "OK")
}

// 离线消息同步
func (ctx *ConnContext) Sync(input defs.Input) {
	var sync defs.SyncInput
	err := json.Unmarshal([]byte(input.Data), &sync)
	if err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	seq, _ := strconv.ParseInt(sync.Seq, 10, 64)

	messageList, err := service.MessageService.ListByUserIdAndSeq(ctx.AppId, ctx.UserId, seq)
	var syncOutput defs.SyncOutput
	if err == nil {
		messageItems := make([]defs.MessageItem, 0, 5)
		for _, v := range *messageList {
			var messageItem defs.MessageItem
			messageItem.SenderId = strconv.FormatInt(v.SenderId, 10)
			messageItem.ReceiverId = strconv.FormatInt(v.ReceiverId, 10)
			messageItem.SendTime = util.FormatDatetime(v.SendTime, util.YYYYMMDDHHMMSS)
			messageItem.Type = defs.MessageType(v.Type)
			messageItem.Content = v.Content
			messageItem.Seq = strconv.FormatInt(v.Seq, 10)
			messageItems = append(messageItems, messageItem)
		}
		syncOutput = defs.SyncOutput{Messages: messageItems}
	}
	ctx.Output(defs.PackageType_SYNC, input.RequestId, err, &syncOutput)
}

func (ctx *ConnContext) Heartbeat(input defs.Input) {
	ctx.Output(defs.PackageType_HEARTBEAT, input.RequestId, nil, "PONG")
	log.Print("device_id：", ctx.DeviceId, " PING")
}

// 消息回执
func (ctx *ConnContext) MessageACK(input defs.Input) {
	var messageACK defs.MessageACK
	err := json.Unmarshal([]byte(input.Data), &messageACK)
	if err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	err = service.DeviceAckService.Update(ctx.DeviceId, messageACK.DeviceAck)
	if err != nil {
		log.Print(err)
	}
	ctx.Output(defs.PackageType_MESSAGE_ACK, input.RequestId, err, "OK")
}

func (ctx *ConnContext) SendToUser(input defs.Input) {
	var message defs.SendMessage
	err := json.Unmarshal([]byte(input.Data), &message)
	if err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	appId, _ := strconv.ParseInt(message.AppId, 10, 64)
	senderId, _ := strconv.ParseInt(message.SenderId, 10, 64)
	deviceId, _ := strconv.ParseInt(message.DeviceId, 10, 64)
	sender := defs.Sender{
		AppId:      appId,
		SenderType: defs.SenderType_ST_SYSTEM,
		SenderId:   senderId,
		DeviceId:   deviceId,
	}
	receiverId, _ := strconv.ParseInt(message.ReceiverId, 10, 64)
	messageReq := defs.SendMessageReq{
		ReceiverType:   message.ReceiverType,
		ReceiverId:     receiverId,
		MessageType:    message.MessageType,
		MessageContent: message.MessageContent,
		ToUserIds:      message.ToUserIds,
		IsPersist:      true,
	}
	err = service.MessageService.SendToFriend(input.RequestId, sender, messageReq)
	if err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
}

func (ctx *ConnContext) Output(pt defs.PackageType, requestId int64, err error, data interface{}) {
	var output = defs.Output{
		Type:      pt,
		RequestId: requestId,
		Data:      data,
	}
	if err != nil {
		output.Code = 1
		output.Message = err.Error()
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
		clear(ctx.DeviceId)
		_ = service.DeviceService.Offline(ctx.AppId, ctx.UserId, ctx.DeviceId)
	}
}
