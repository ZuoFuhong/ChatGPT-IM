package service

import (
	"go-IM/internal/logic/dao"
	"go-IM/internal/logic/model"
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

// TODO: 消息发送
