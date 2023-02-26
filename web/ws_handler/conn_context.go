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
	DeviceId int64
	UserId   int64
}

func NewConnContext(conn *websocket.Conn) *ConnContext {
	return &ConnContext{
		Conn: conn,
	}
}

func (ctx *ConnContext) DoConn() {
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
	input := new(defs.Input)
	if err := json.Unmarshal(data, input); err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	switch input.Type {
	case defs.PackagetypeSignin:
		ctx.SignIn(input)
	case defs.PackagetypeSync:
		ctx.Sync(input)
	case defs.PackagetypeHeartbeat:
		ctx.Heartbeat(input)
	case defs.PackagetypeMessageAck:
		ctx.MessageACK(input)
	case defs.PackagetypeRtUser:
		ctx.SendToUser(input)
	}
}

// SignIn 登录鉴权
func (ctx *ConnContext) SignIn(input *defs.Input) {
	signIn := new(defs.SignIn)
	if err := json.Unmarshal([]byte(input.Data), &signIn); err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	userId, _ := strconv.ParseInt(signIn.UserId, 10, 64)
	deviceId, _ := strconv.ParseInt(signIn.DeviceId, 10, 64)
	if err := service.AuthService.SignIn(userId, deviceId, signIn.Token, "", 0); err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	ctx.UserId = userId
	ctx.DeviceId = deviceId

	// 断开这个设备之前的连接
	preCtx := load(ctx.DeviceId)
	if preCtx != nil {
		preCtx.DeviceId = PreConn
	}
	store(ctx.DeviceId, ctx)
	ctx.Output(defs.PackagetypeSignin, input.RequestId, nil, "OK")
}

// Sync 离线消息同步
func (ctx *ConnContext) Sync(input *defs.Input) {
	var sync defs.SyncInput
	if err := json.Unmarshal([]byte(input.Data), &sync); err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	seq, _ := strconv.ParseInt(sync.Seq, 10, 64)
	msgList, err := service.MessageService.ListDeviceMessageBySeq(ctx.UserId, ctx.DeviceId, seq)
	var syncOutput defs.SyncOutput
	if err == nil {
		messageItems := make([]defs.MessageItem, 0)
		for _, v := range msgList {
			var messageItem defs.MessageItem
			messageItem.SenderId = strconv.FormatInt(v.SenderId, 10)
			messageItem.ReceiverId = strconv.FormatInt(v.ReceiverId, 10)
			messageItem.SendTime = util.FormatDatetime(v.SendTime)
			messageItem.Type = defs.MessageType(v.Type)
			messageItem.Content = v.Content
			messageItem.Seq = strconv.FormatInt(v.Seq, 10)
			messageItems = append(messageItems, messageItem)
		}
		syncOutput = defs.SyncOutput{Messages: messageItems}
	}
	ctx.Output(defs.PackagetypeSync, input.RequestId, err, &syncOutput)
}

func (ctx *ConnContext) Heartbeat(input *defs.Input) {
	log.Printf("device_id:%d %s", ctx.DeviceId, input.Data)
	ctx.Output(defs.PackagetypeHeartbeat, input.RequestId, nil, "PONG")
}

// MessageACK 消息回执
func (ctx *ConnContext) MessageACK(input *defs.Input) {
	var messageACK defs.MessageACK
	if err := json.Unmarshal([]byte(input.Data), &messageACK); err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	err := service.DeviceAckService.Update(ctx.DeviceId, messageACK.DeviceAck)
	if err != nil {
		log.Print(err)
	}
	ctx.Output(defs.PackagetypeMessageAck, input.RequestId, err, "OK")
}

func (ctx *ConnContext) SendToUser(input *defs.Input) {
	var message defs.SendMessage
	if err := json.Unmarshal([]byte(input.Data), &message); err != nil {
		log.Print(err)
		ctx.Release()
		return
	}
	senderId, _ := strconv.ParseInt(message.SenderId, 10, 64)
	deviceId, _ := strconv.ParseInt(message.DeviceId, 10, 64)
	sender := &defs.Sender{
		SenderType: defs.SendertypeStSystem,
		SenderId:   senderId,
		DeviceId:   deviceId,
	}
	receiverId, _ := strconv.ParseInt(message.ReceiverId, 10, 64)
	messageReq := &defs.SendMessageReq{
		ReceiverType:   message.ReceiverType,
		ReceiverId:     receiverId,
		MessageType:    message.MessageType,
		MessageContent: message.MessageContent,
		ToUserIds:      message.ToUserIds,
		IsPersist:      true,
	}
	// 1.消息持久化 2.查询用户在线设备 3.消费发送给用户设备
	if err := service.MessageService.Send(input.RequestId, sender, messageReq); err != nil {
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

// Release 释放 TCP 连接
func (ctx *ConnContext) Release() {
	e := ctx.Conn.Close()
	if e != nil {
		log.Print(e)
	}
	// 设备下线
	if ctx.DeviceId != PreConn {
		clear(ctx.DeviceId)
		_ = service.DeviceService.Offline(ctx.DeviceId)
	}
}
