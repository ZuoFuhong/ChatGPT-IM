package dao

import "go-IM/internal/logic/db"

type deviceAckDao struct{}

var DeviceAckDao = new(deviceAckDao)

// 添加设备同步序列号记录
func (*deviceAckDao) Add(deviceId int64, ack int64) error {
	_, err := db.Cli.Exec("insert into device_ack(device_id,ack) values(?,?)", deviceId, ack)
	return err
}

// 获取设备同步序列号
func (*deviceAckDao) Get(deviceId int64) (int64, error) {
	row := db.Cli.QueryRow("select ack from device_ack where device_id = ?", deviceId)
	var ack int64
	err := row.Scan(&ack)
	return ack, err
}

// 更新设备同步序列号
func (*deviceAckDao) Update(deviceId, ack int64) error {
	_, e := db.Cli.Exec("UPDATE device_ack SET ack = ? WHERE device_id = ?", ack, deviceId)
	return e
}

// 获取用户最大的同步序列号
func (*deviceAckDao) GetMaxByUserId(appId, userId int64) (int64, error) {
	row := db.Cli.QueryRow(`
		SELECT max(a.ack)
		FROM device d
		LEFT JOIN device_ack a ON d.device_id = a.device_id
		WHERE d.app_id = ?
		AND d.user_id = ?`, appId, userId)
	var ack int64
	err := row.Scan(&ack)
	return ack, err
}
