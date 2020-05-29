package service

import (
	"errors"
	"go-IM/internal/logic/dao"
	"go-IM/internal/logic/model"
	"go-IM/pkg/defs"
	"log"
	"strconv"
	"strings"
	"time"
)

type messageService struct{}

var MessageService = new(messageService)

// Add 添加消息
func (*messageService) Add(message model.Message) error {
	return dao.MessageDao.Add(message)
}

// 查询消息
func (*messageService) ListByUserIdAndSeq(appId, userId, seq int64) (*[]model.Message, error) {
	var err error
	if seq == 0 {
		seq, err = DeviceAckService.GetMaxByUserId(appId, userId)
		if err != nil {
			return nil, err
		}
	}
	messages, err := dao.MessageDao.ListBySeq(appId, model.MessageObjectTypeUser, userId, seq)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// Send 消息发送
func (s *messageService) Send(requestId int64, sender defs.Sender, req defs.SendMessageReq) error {
	switch req.ReceiverType {
	case defs.ReceiverType_RT_USER:
		if sender.SenderType == defs.SenderType_ST_USER {
			err := MessageService.SendToFriend(requestId, sender, req)
			if err != nil {
				return err
			}
		} else {
			err := MessageService.SendToUser(requestId, sender, req.ReceiverId, 0, req)
			if err != nil {
				return err
			}
		}
	case defs.ReceiverType_RT_NORMAL_GROUP:
		err := MessageService.SendToGroup(requestId, sender, req)
		if err != nil {
			return err
		}
	case defs.ReceiverType_RT_LARGE_GROUP:
		err := MessageService.SendToChatRoom(requestId, sender, req)
		if err != nil {
			return err
		}
	}
	return nil
}

// SendToUser 消息发送至用户
func (*messageService) SendToFriend(requestId int64, sender defs.Sender, req defs.SendMessageReq) error {
	// 发给发送者
	err := MessageService.SendToUser(requestId, sender, sender.SenderId, 0, req)
	if err != nil {
		return err
	}

	// 发给接收者
	err = MessageService.SendToUser(requestId, sender, req.ReceiverId, 0, req)
	if err != nil {
		return err
	}

	return nil
}

// 消息发送至群组（使用写扩散）
func (*messageService) SendToGroup(requestId int64, sender defs.Sender, req defs.SendMessageReq) error {
	users := GroupUserService.GetUsers(sender.AppId, req.ReceiverId)

	if sender.SenderType == defs.SenderType_ST_USER && !IsInGroup(users, sender.SenderId) {
		log.Print(sender.SenderId, req.ReceiverId, "不在群组内")
		return errors.New("not in group")
	}

	// 将消息发送给群组用户，使用写扩散
	for _, user := range users {
		err := MessageService.SendToUser(requestId, sender, user.UserId, 0, req)
		if err != nil {
			return err
		}
	}
	return nil
}

func IsInGroup(users []model.GroupUser, userId int64) bool {
	for i := range users {
		if users[i].UserId == userId {
			return true
		}
	}
	return false
}

// 消息发送至聊天室（读扩散）
func (*messageService) SendToChatRoom(requestId int64, sender defs.Sender, req defs.SendMessageReq) error {
	users := GroupUserService.GetUsers(sender.AppId, req.ReceiverId)
	if !IsInGroup(users, sender.SenderId) {
		return errors.New("Not in group")
	}
	var seq int64 = 0
	if req.IsPersist {
		seq, err := SeqService.GetGroupNext(sender.AppId, req.ReceiverId)
		if err != nil {
			return err
		}
		message := model.Message{
			AppId:          sender.AppId,
			ObjectType:     model.MessageObjectTypeGroup,
			ObjectId:       req.ReceiverId,
			RequestId:      requestId,
			SenderType:     int32(sender.SenderType),
			SenderId:       sender.SenderId,
			SenderDeviceId: sender.DeviceId,
			ReceiverType:   int32(req.ReceiverType),
			ReceiverId:     req.ReceiverId,
			ToUserIds:      FormatUserIds(req.ToUserIds),
			Type:           req.MessageType,
			Content:        string(req.MessageContent),
			Seq:            seq,
			SendTime:       time.Now(),
			Status:         int32(defs.MessageStatus_MS_NORMAL),
		}
		err = MessageService.Add(message)
		if err != nil {
			return err
		}
	}

	// 将消息发送给群组用户，使用读扩散
	req.IsPersist = false
	for _, v := range users {
		err := MessageService.SendToUser(requestId, sender, v.UserId, seq, req)
		if err != nil {
			return err
		}
	}
	return nil
}

// 将消息持久化到数据库,并且消息发送至用户
func (*messageService) SendToUser(requestId int64, sender defs.Sender, toUserId int64, roomSeq int64, req defs.SendMessageReq) error {
	log.Print("message_store_send_to_user", " app_id：", sender.AppId, " to_user_id：", toUserId)

	seq, err := SeqService.GetUserNext(sender.AppId, toUserId)
	if err != nil {
		return err
	}
	message := model.Message{
		AppId:          sender.AppId,
		ObjectType:     model.MessageObjectTypeUser,
		ObjectId:       toUserId,
		RequestId:      requestId,
		SenderType:     int32(sender.SenderType),
		SenderId:       sender.SenderId,
		SenderDeviceId: sender.DeviceId,
		ReceiverType:   int32(req.ReceiverType),
		ReceiverId:     req.ReceiverId,
		ToUserIds:      FormatUserIds(req.ToUserIds),
		Type:           req.MessageType,
		Content:        string(req.MessageContent),
		Seq:            seq,
		SendTime:       time.Now(),
		Status:         int32(defs.MessageStatus_MS_NORMAL),
	}
	if req.IsPersist {
		err := MessageService.Add(message)
		if err != nil {
			return err
		}
	}
	// 查询用户在线设备
	devices, err := DeviceService.ListOnlineByUserId(sender.AppId, toUserId)
	if err != nil {
		return err
	}
	for _, v := range devices {
		if sender.DeviceId == v.DeviceId {
			continue
		}
		err := MessageService.SendToDevice(v, message)
		if err != nil {
			return err
		}
	}
	return nil
}

// 将消息发送给设备
func (*messageService) SendToDevice(device model.Device, message model.Message) error {
	if device.Status == model.DeviceOnLine {
		// todo: load(deviceId)
	}
	return nil
}

func FormatUserIds(userId []int64) string {
	build := strings.Builder{}
	for i, v := range userId {
		build.WriteString(strconv.FormatInt(v, 10))
		if i != len(userId)-1 {
			build.WriteString(",")
		}
	}
	return build.String()
}
