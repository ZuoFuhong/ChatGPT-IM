package service

import "go-IM/internal/logic/dao"

type seqService struct{}

var SeqService = new(seqService)

// 获取下一个序列号
func (*seqService) GetUserNext(appId, userId int64) (int64, error) {
	return dao.MessageDao.GetMaxByObjectId(appId, 1, userId)
}

// 获取下一个序列号
func (*seqService) GetGroupNext(appId, groupId int64) (int64, error) {
	return dao.MessageDao.GetMaxByObjectId(appId, 2, groupId)
}
