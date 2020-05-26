package service

import "go-IM/internal/logic/dao"

type deviceAckService struct{}

var DeviceAckService = new(deviceAckService)

// 更新设备同步序列号
func (*deviceAckService) Update(deviceId, ack int64) error {
	return dao.DeviceAckDao.Update(deviceId, ack)
}

// 获取用户最大的同步序列号
func (*deviceAckService) GetMaxByUserId(appId, userId int64) (int64, error) {
	return dao.DeviceAckDao.GetMaxByUserId(appId, userId)
}
