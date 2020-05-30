package service

import (
	"go-IM/logic/dao"
	"go-IM/logic/model"
	"go-IM/pkg/errs"
	"go-IM/pkg/util"
)

const (
	DeviceOnline  = 1
	DeviceOffline = 0
)

type deviceService struct{}

var DeviceService = new(deviceService)

func (*deviceService) Register(device *model.Device) (int64, error) {
	app, err := AppService.Get(device.AppId)
	if err != nil {
		panic(err)
	}
	if app == nil {
		panic(errs.NewHttpErr(errs.App, "App not found"))
	}

	// 唯一的设备ID
	sf, _ := util.NewSnowflake(0, 0)
	deviceId := sf.NextVal()

	device.DeviceId = deviceId
	err = dao.DeviceDao.Add(device)
	if err != nil {
		panic(err)
	}
	err = dao.DeviceAckDao.Add(device.DeviceId, 0)
	if err != nil {
		panic(err)
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
