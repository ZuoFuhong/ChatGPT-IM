package model

import (
	"encoding/json"
	"fmt"
)

const (
	DeviceOnLine  = 1 // 设备在线
	DeviceOffLine = 0 // 设备离线
)

const (
	Android  = 1
	Ios      = 2
	Windowds = 3
	MacOS    = 4
	Web      = 5
)

// Device 设备
type Device struct {
	Id       int64  `json:"id"`        // 设备ID
	UserId   int64  `json:"user_id"`   // 用户ID
	Type     int32  `json:"type"`      // 设备类型 1:Android；2：IOS；3：Windows; 4：MacOS；5：Web
	Status   int32  `json:"status"`    // 在线状态，0：不在线；1：在线
	ConnAddr string `json:"conn_addr"` // 连接层服务层地址
	ConnFd   int64  `json:"conn_fd"`   // TCP连接对应的文件描述符
	Ack      int64  `json:"ack"`       // 同步序列号
	Ctime    int64  `json:"ctime"`     // 创建时间
	Utime    int64  `json:"utime"`     // 更新时间
}

func (*Device) BucketName() string {
	return "device"
}

// StoreDevice 存储设备信息
func StoreDevice(dm *Device) error {
	bytes, err := json.Marshal(dm)
	if err != nil {
		return err
	}
	return WriteToDB(dm.BucketName(), fmt.Sprint(dm.Id), bytes)
}

// LoadDevice 读取设备信息
func LoadDevice(deviceId int64) (*Device, error) {
	dm := &Device{}
	bytes, err := ReadFromDB(dm.BucketName(), fmt.Sprint(deviceId))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, dm); err != nil {
		return nil, err
	}
	return dm, nil
}
