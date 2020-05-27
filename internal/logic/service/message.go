package service

import (
	"errors"
	"go-IM/internal/logic/dao"
	"go-IM/internal/logic/model"
	"go-IM/pkg/defs"
	"log"
)

type messageService struct{}

var MessageService = new(messageService)

// Add 添加消息
func (*messageService) Add(message model.Message) error {
	return dao.MessageDao.Add("message", message)
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
func (s *messageService) Send(sender defs.Sender, req defs.SendMessageReq) error {
	switch req.ReceiverType {
	case defs.ReceiverType_RT_USER:
		if sender.SenderType == defs.SenderType_ST_USER {
			err := MessageService.SendToFriend(sender, req)
			if err != nil {
				return err
			}
		} else {
			err := MessageService.SendToUser(sender, req.ReceiverId, 0, req)
			if err != nil {
				return err
			}
		}
	case defs.ReceiverType_RT_NORMAL_GROUP:
		err := MessageService.SendToGroup(sender, req)
		if err != nil {
			return err
		}
	case defs.ReceiverType_RT_LARGE_GROUP:
		err := MessageService.SendToChatRoom(sender, req)
		if err != nil {
			return err
		}
	}

	return nil
}

// SendToUser 消息发送至用户
func (*messageService) SendToFriend(sender defs.Sender, req defs.SendMessageReq) error {
	// 发给发送者
	err := MessageService.SendToUser(sender, sender.SenderId, 0, req)
	if err != nil {
		return err
	}

	// 发给接收者
	err = MessageService.SendToUser(sender, req.ReceiverId, 0, req)
	if err != nil {
		return err
	}

	return nil
}

// SendToGroup 消息发送至群组（使用写扩散）
func (*messageService) SendToGroup(sender defs.Sender, req defs.SendMessageReq) error {
	users, err := GroupUserService.GetUsers(sender.AppId, req.ReceiverId)
	if err != nil {
		return err
	}

	if sender.SenderType == defs.SenderType_ST_USER && !IsInGroup(users, sender.SenderId) {
		log.Print(sender.SenderId, req.ReceiverId, "不在群组内")
		return errors.New("not in group")
	}

	// 将消息发送给群组用户，使用写扩散
	for _, user := range users {
		err = MessageService.SendToUser(sender, user.UserId, 0, req)
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

// SendToChatRoom 消息发送至聊天室（读扩散）
func (*messageService) SendToChatRoom(sender defs.Sender, req defs.SendMessageReq) error {
	// todo:

	return nil
}

// 将消息持久化到数据库,并且消息发送至用户
func (*messageService) SendToUser(sender defs.Sender, toUserId int64, roomSeq int64, req defs.SendMessageReq) error {

	// todo:

	return nil
}
