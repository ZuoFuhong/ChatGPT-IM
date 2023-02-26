package service

import (
	"errors"
	"go-IM/consts"
	"go-IM/pkg/util"
	"time"
)

type authService struct{}

var AuthService = new(authService)

// SignIn 长连接登录
func (auth *authService) SignIn(userId, deviceId int64, token string, connAddr string, conFd int64) error {
	if err := auth.verifyToken(userId, deviceId, token); err != nil {
		return err
	}
	// 标记用户在设备上登录
	return DeviceService.Online(deviceId, connAddr, conFd)
}

func (*authService) verifyToken(userId, deviceId int64, token string) error {
	tInfo, err := util.DecryptToken(token, consts.PrivateKey)
	if err != nil {
		return err
	}
	if !(tInfo.UserId == userId && tInfo.DeviceId == deviceId) {
		return errors.New("unauthorized")
	}
	if tInfo.Expire < time.Now().Unix() {
		return errors.New("unauthorized")
	}
	return nil
}
