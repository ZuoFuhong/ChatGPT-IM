package service

import (
	model2 "ChatGPT-IM/backend/logic/model"
	"ChatGPT-IM/backend/pkg/tinyid"
	"errors"
	"time"
)

type deviceService struct{}

var DeviceService = new(deviceService)

// Register 设备注册
func (*deviceService) Register(dm *model2.Device) (int64, error) {
	// 查询用户设备列表
	um, err := model2.LoadUser(dm.UserId)
	if err != nil {
		return 0, err
	}
	// 检查重复注册
	for _, deId := range um.Devices {
		cdm, err := model2.LoadDevice(deId)
		if err != nil {
			continue
		}
		if cdm.Type == dm.Type {
			return cdm.Id, nil
		}
	}
	now := time.Now()
	// 唯一的设备ID
	deviceId := tinyid.NextId()
	dm.Id = deviceId
	dm.Ctime = now.UnixMilli()
	dm.Utime = now.UnixMilli()
	if err := model2.StoreDevice(dm); err != nil {
		return 0, err
	}
	// 更新设备列表
	um.Devices = append(um.Devices, deviceId)
	if err := model2.StoreUser(um); err != nil {
		return 0, err
	}
	return deviceId, nil
}

// GetUserWebDevice 查询用户 Web 设备
func (*deviceService) GetUserWebDevice(userId int64) (*model2.Device, error) {
	um, err := model2.LoadUser(userId)
	if err != nil {
		return nil, err
	}
	for _, deId := range um.Devices {
		cdm, err := model2.LoadDevice(deId)
		if err != nil {
			continue
		}
		if cdm.Type == model2.Web {
			return cdm, nil
		}
	}
	return nil, errors.New("has not register web device")
}

// ListOnlineByUid 获取用户的所有在线设备
func (*deviceService) ListOnlineByUid(userId int64) ([]*model2.Device, error) {
	user, err := model2.LoadUser(userId)
	if err != nil {
		return nil, err
	}
	dmList := make([]*model2.Device, 0)
	for _, deviceId := range user.Devices {
		dm, err := model2.LoadDevice(deviceId)
		if err != nil {
			return nil, err
		}
		if dm.Status == model2.DeviceOnLine {
			dmList = append(dmList, dm)
		}
	}
	return dmList, nil
}

// Online 设备上线
func (*deviceService) Online(deviceId int64, connAddr string, connFd int64) error {
	dm, err := model2.LoadDevice(deviceId)
	if err != nil {
		return err
	}
	dm.Status = model2.DeviceOnLine
	dm.ConnAddr = connAddr
	dm.ConnFd = connFd
	return model2.StoreDevice(dm)
}

// Offline 设备离线
func (*deviceService) Offline(deviceId int64) error {
	dm, err := model2.LoadDevice(deviceId)
	if err != nil {
		return err
	}
	dm.Status = model2.DeviceOffLine
	return model2.StoreDevice(dm)
}
