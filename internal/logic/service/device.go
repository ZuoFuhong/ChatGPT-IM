package service

import (
	"errors"
	"go-IM/internal/logic/dao"
	"go-IM/internal/logic/model"
)

const (
	DeviceOnline  = 1
	DeviceOffline = 0
)

type deviceService struct{}

var DeviceService = new(deviceService)

func (*deviceService) Register(device model.Device) (int64, error) {
	app, err := AppService.Get(device.AppId)
	if err != nil {
		return 0, err
	}
	if app == nil {
		return 0, errors.New("app not found")
	}

	// todo: 自增ID
	deviceId := int64(0)

	device.DeviceId = deviceId
	err = dao.DeviceDao.Add(device)
	if err != nil {
		return 0, err
	}
	err = dao.DeviceAckDao.Add(device.DeviceId, 0)
	if err != nil {
		return 0, err
	}
	return deviceId, nil
}

// 获取用户的所有在线设备
func (*deviceService) ListOnlineByUserId(appId, userId int64) ([]model.Device, error) {
	devices, err := dao.DeviceDao.ListOnlineByUserId(appId, userId)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

// 设备上线
func (*deviceService) Online(appId, deviceId, userId int64, connAddr string, connFd int64) error {
	return dao.DeviceDao.UpdateUserIdAndStatus(deviceId, userId, DeviceOnline, connAddr, connFd)
}

// 设备离线
func (*deviceService) Offline(appId, userId, deviceId int64) error {
	return dao.DeviceDao.UpdateStatus(deviceId, DeviceOffline)
}
