package service

import (
	"go-IM/logic/model"
	"time"
)

type deviceAckService struct{}

var DeviceAckService = new(deviceAckService)

// Update 更新设备同步序列号
func (*deviceAckService) Update(deviceId, ack int64) error {
	dm, err := model.LoadDevice(deviceId)
	if err != nil {
		return err
	}
	dm.Ack = ack
	dm.Utime = time.Now().UnixMilli()
	return model.StoreDevice(dm)
}

// GetDeviceMaxSeq 获取设备最大的同步序列号
func (*deviceAckService) GetDeviceMaxSeq(deviceId int64) (int64, error) {
	dm, err := model.LoadDevice(deviceId)
	if err != nil {
		return 0, err
	}
	return dm.Ack, nil
}
