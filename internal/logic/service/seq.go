package service

type seqService struct{}

var SeqService = new(seqService)

// 获取下一个序列号
func (*seqService) GetUserNext(appId, userId int64) (int64, error) {

	// TODO: 使用缓存自增

	return 0, nil
}

// 获取下一个序列号
func (*seqService) GetGroupNext(appId, groupId int64) (int64, error) {

	return 0, nil
}
