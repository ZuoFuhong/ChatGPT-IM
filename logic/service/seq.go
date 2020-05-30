package service

import (
	"go-IM/logic/dao"
	"go-IM/logic/model"
)

type seqService struct{}

var SeqService = new(seqService)

// 获取下一个序列号
func (*seqService) GetUserNext(appId, userId int64) (int64, error) {
	maxSeq, e := dao.MessageDao.GetMaxByObjectId(appId, model.MessageObjectTypeUser, userId)
	return maxSeq + 1, e
}

// 获取下一个序列号
func (*seqService) GetGroupNext(appId, groupId int64) (int64, error) {
	maxSeq, e := dao.MessageDao.GetMaxByObjectId(appId, model.MessageObjectTypeGroup, groupId)
	return maxSeq + 1, e
}
