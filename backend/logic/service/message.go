package service

import (
	"ChatGPT-IM/backend/consts"
	model2 "ChatGPT-IM/backend/logic/model"
	"ChatGPT-IM/backend/pkg/defs"
	"ChatGPT-IM/backend/pkg/tinyid"
	"ChatGPT-IM/backend/pkg/util"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

type messageService struct {
	PushChan chan defs.MessageItem
}

var MessageService = new(messageService)

// ListDeviceMessageBySeq 查询设备落后的消息
func (*messageService) ListDeviceMessageBySeq(uid, deviceId, seq int64) ([]*model2.Message, error) {
	var err error
	if seq == 0 {
		seq, err = DeviceAckService.GetDeviceMaxSeq(deviceId)
		if err != nil {
			return nil, err
		}
	}
	return model2.ListBySeq(model2.MessageObjectTypeUser, uid, seq)
}

// Send 消息发送
func (s *messageService) Send(requestId int64, sender *defs.Sender, req *defs.SendMessageReq) error {
	switch req.ReceiverType {
	case defs.ReceivertypeRtUser:
		// 发给发送者
		if err := s.SendToUser(requestId, sender, sender.SenderId, req); err != nil {
			return err
		}
		// 发给接收者
		if err := s.SendToUser(requestId, sender, req.ReceiverId, req); err != nil {
			return err
		}
	case defs.ReceivertypeRtNormalGroup:
		err := s.SendToGroup(requestId, sender, req)
		if err != nil {
			return err
		}
	case defs.ReceivertypeRtLargeGroup:
		if err := s.SendToChatRoom(requestId, sender, req); err != nil {
			return err
		}
	}
	return nil
}

// SendToUser 将消息持久化到数据库,并且消息发送至用户
func (s *messageService) SendToUser(requestId int64, sender *defs.Sender, toUserId int64, req *defs.SendMessageReq) error {
	log.Print("message_store_send_to_user", " to_user_id：", toUserId)

	// 表示向机器人提问
	if toUserId == consts.RobotUid && req.MessageType == defs.MessageText {
		if err := s.RobotRelay(sender, req); err != nil {
			log.Printf("RobotRelay failed, senderId=%v, toUserId=%v, err:%v\n", sender.SenderId, toUserId, err)
		}
		return nil
	}

	// 获取下一个序列号
	seq := tinyid.NextId()
	msg := &model2.Message{
		ObjectType:     model2.MessageObjectTypeUser,
		ObjectId:       toUserId,
		RequestId:      requestId,
		SenderType:     int32(sender.SenderType),
		SenderId:       sender.SenderId,
		SenderDeviceId: sender.DeviceId,
		ReceiverType:   int32(req.ReceiverType),
		ReceiverId:     req.ReceiverId,
		ToUserIds:      FormatUserIds(req.ToUserIds),
		Type:           int(req.MessageType),
		Content:        req.MessageContent,
		Seq:            seq,
		SendTime:       time.Now().UnixMilli(),
		Status:         int32(defs.MessagestatusMsNormal),
	}
	// 消息持久化
	if req.IsPersist {
		if err := model2.StoreMessage(msg); err != nil {
			return err
		}
	}
	// 查询用户在线设备
	devices, err := DeviceService.ListOnlineByUid(toUserId)
	if err != nil {
		return err
	}
	for _, v := range devices {
		err := s.SendToDevice(v, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

// RobotRelay 机器人智能回复
func (s *messageService) RobotRelay(sender *defs.Sender, req *defs.SendMessageReq) error {
	// 代理机器人回答
	answer, err := ProxyRobotPost(req.MessageContent)
	if err != nil {
		answer = err.Error()
	}
	// 获取下一个序列号
	seq := tinyid.NextId()
	msg := &model2.Message{
		ObjectType:     model2.MessageObjectTypeUser,
		ObjectId:       sender.SenderId,
		RequestId:      0,
		SenderType:     int32(defs.SendertypeStSystem),
		SenderId:       consts.RobotUid,
		SenderDeviceId: 0,
		ReceiverType:   int32(req.ReceiverType),
		ReceiverId:     sender.SenderId,
		ToUserIds:      FormatUserIds(req.ToUserIds),
		Type:           int(req.MessageType),
		Content:        answer,
		Seq:            seq,
		SendTime:       time.Now().UnixMilli(),
		Status:         int32(defs.MessagestatusMsNormal),
	}
	// 消息持久化
	if err := model2.StoreMessage(msg); err != nil {
		return err
	}
	// 查询用户在线设备
	devices, err := DeviceService.ListOnlineByUid(sender.SenderId)
	if err != nil {
		return err
	}
	for _, v := range devices {
		err := s.SendToDevice(v, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

// SendToDevice 将消息发送给设备
func (s *messageService) SendToDevice(device *model2.Device, msg *model2.Message) error {
	// 目前仅支持推送 Web
	if device.Status == model2.DeviceOnLine && device.Type == model2.Web {
		var messageItem defs.MessageItem
		messageItem.SenderId = strconv.FormatInt(msg.SenderId, 10)
		messageItem.ReceiverId = strconv.FormatInt(msg.ReceiverId, 10)
		messageItem.ReceiverDeviceId = device.Id
		messageItem.SendTime = util.FormatDatetime(msg.SendTime)
		messageItem.Type = defs.MessageType(msg.Type)
		messageItem.Content = msg.Content
		messageItem.Seq = strconv.FormatInt(msg.Seq, 10)
		// 使用 chan 交互
		s.PushChan <- messageItem
	}
	return nil
}

func FormatUserIds(userId []string) string {
	build := strings.Builder{}
	for i, v := range userId {
		build.WriteString(v)
		if i != len(userId)-1 {
			build.WriteString(",")
		}
	}
	return build.String()
}

// SendToGroup 消息发送至群组（使用写扩散）
func (s *messageService) SendToGroup(requestId int64, sender *defs.Sender, req *defs.SendMessageReq) error {
	// Todo: 当前没有实现查询群组成员
	users := make([]*model2.GroupUser, 0)
	if sender.SenderType == defs.SendertypeStUser && !IsInGroup(users, sender.SenderId) {
		log.Print(sender.SenderId, req.ReceiverId, "不在群组内")
		return errors.New("not in group")
	}

	// 将消息发送给群组用户，使用写扩散
	for _, user := range users {
		err := s.SendToUser(requestId, sender, user.UserId, req)
		if err != nil {
			return err
		}
	}
	return nil
}

// IsInGroup 检查群组成员
func IsInGroup(users []*model2.GroupUser, userId int64) bool {
	for i := range users {
		if users[i].UserId == userId {
			return true
		}
	}
	return false
}

// SendToChatRoom 消息发送至聊天室（读扩散）
func (s *messageService) SendToChatRoom(requestId int64, sender *defs.Sender, req *defs.SendMessageReq) error {
	// Todo: 当前没有实现查询群组成员
	users := make([]*model2.GroupUser, 0)
	receiverId := req.ReceiverId
	if !IsInGroup(users, sender.SenderId) {
		return errors.New("not in group")
	}
	if req.IsPersist {
		// 获取下一个序列号
		seq := tinyid.NextId()
		msg := &model2.Message{
			ObjectType:     model2.MessageObjectTypeGroup,
			ObjectId:       receiverId,
			RequestId:      requestId,
			SenderType:     int32(sender.SenderType),
			SenderId:       sender.SenderId,
			SenderDeviceId: sender.DeviceId,
			ReceiverType:   int32(req.ReceiverType),
			ReceiverId:     receiverId,
			ToUserIds:      FormatUserIds(req.ToUserIds),
			Type:           int(req.MessageType),
			Content:        req.MessageContent,
			Seq:            seq,
			SendTime:       time.Now().UnixMilli(),
			Status:         int32(defs.MessagestatusMsNormal),
		}
		if err := model2.StoreMessage(msg); err != nil {
			return err
		}
	}
	// 将消息发送给群组用户，使用读扩散
	req.IsPersist = false
	for _, v := range users {
		err := s.SendToUser(requestId, sender, v.UserId, req)
		if err != nil {
			return err
		}
	}
	return nil
}
