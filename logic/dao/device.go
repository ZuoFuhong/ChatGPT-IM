package dao

import (
	"database/sql"
	"go-IM/logic/db"
	"go-IM/logic/model"
)

type deviceDao struct{}

var DeviceDao = new(deviceDao)

// 插入一条设备信息
func (*deviceDao) Add(device *model.Device) error {
	_, err := db.Cli.Exec(`
		INSERT INTO device(device_id,app_id,user_id,type,brand,model,system_version,sdk_version,status,conn_addr,conn_fd) 
		VALUES(?,?,?,?,?,?,?,?,?,?,?)`,
		device.DeviceId, device.AppId, device.UserId, device.Type, device.Brand, device.Model, device.SystemVersion, device.SDKVersion, device.Status, "", 0)
	return err
}

// Get 获取设备
func (*deviceDao) Get(deviceId int64) (*model.Device, error) {
	device := model.Device{
		DeviceId: deviceId,
	}
	row := db.Cli.QueryRow(`
		SELECT app_id,user_id,type,brand,model,system_version,sdk_version,status,conn_addr,conn_fd,create_time,update_time
		FROM device
		WHERE device_id = ?`, deviceId)
	err := row.Scan(&device.AppId, &device.UserId, &device.Type, &device.Brand, &device.Model, &device.SystemVersion, &device.SDKVersion,
		&device.Status, &device.ConnAddr, &device.ConnFd, &device.CreateTime, &device.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &device, err
}

// 查询用户所有的在线设备
func (*deviceDao) ListOnlineByUserId(appId, userId int64) ([]model.Device, error) {
	rows, err := db.Cli.Query(`
		SELECT device_id,type,brand,model,system_version,sdk_version,status,conn_addr,conn_fd,create_time,update_time 
		FROM device 
		WHERE app_id = ? 
		AND user_id = ?
		AND status = ?`,
		appId, userId, model.DeviceOnLine)
	if err != nil {
		return nil, err
	}

	devices := make([]model.Device, 0, 5)
	for rows.Next() {
		device := new(model.Device)
		err = rows.Scan(&device.DeviceId, &device.Type, &device.Brand, &device.Model, &device.SystemVersion, &device.SDKVersion,
			&device.Status, &device.ConnAddr, &device.ConnFd, &device.CreateTime, &device.UpdateTime)
		if err != nil {
			return nil, err
		}
		devices = append(devices, *device)
	}
	return devices, nil
}

// 更新设备绑定用户和设备在线状态
func (*deviceDao) UpdateUserIdAndStatus(deviceId, userId int64, status int, connAddr string, connFd int64) error {
	_, e := db.Cli.Exec(`
		UPDATE device 
		SET user_id = ?, status = ?, conn_addr = ?, conn_fd = ? 
		WHERE device_id = ?`, userId, status, connAddr, connFd, deviceId)
	return e
}

// 更新设备的在线状态
func (*deviceDao) UpdateStatus(deviceId int64, status int) error {
	_, e := db.Cli.Exec("UPDATE device SET status = ? WHERE device_id = ?", status, deviceId)
	return e
}

// 升级设备
func (*deviceDao) Upgrade(deviceId int64, systemVersion, sdkVersion string) error {
	_, err := db.Cli.Exec("update device set system_version = ?,sdk_version = ? where device_id = ? ",
		systemVersion, sdkVersion, deviceId)
	return err
}

// 测试使用
func (*deviceDao) GetDeviceByParams(userId int64, brand string) (int64, error) {
	row := db.Cli.QueryRow("SELECT device_id FROM device WHERE user_id = ? AND brand = ?", userId, brand)
	var deviceId int64
	err := row.Scan(&deviceId)
	return deviceId, err
}
