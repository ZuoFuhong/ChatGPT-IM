package service

import (
	"errors"
	"go-IM/pkg/util"
	"time"
)

type authService struct{}

var AuthService = new(authService)

// 长连接登录
func (auth *authService) SignIn(appId, userId, deviceId int64, token string, connAddr string, conFd int64) error {
	err := auth.verifyToken(appId, userId, deviceId, token)
	if err != nil {
		return err
	}

	// 标记用户在设备上登录
	err = DeviceService.Online(appId, deviceId, userId, connAddr, conFd)
	return err
}

func (*authService) verifyToken(appId, userId, deviceId int64, token string) error {
	app, e := AppService.Get(appId)
	if e != nil {
		return e
	}
	if app == nil {
		return errors.New("app not found")
	}
	info, e := util.DecryptToken(token, app.PrivateKey)
	if e != nil {
		return e
	}
	if !(info.AppId == appId && info.UserId == userId && info.DeviceId == deviceId) {
		return errors.New("unauthorized")
	}
	if info.Expire < time.Now().Unix() {
		return errors.New("unauthorized")
	}
	return nil
}
