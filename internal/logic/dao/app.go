package dao

import (
	"go-IM/internal/logic/db"
	"go-IM/internal/logic/model"
)

type appDao struct{}

var AppDao = new(appDao)

// 获取APP信息
func (*appDao) Get(appId int64) (*model.App, error) {
	var app model.App
	err := db.Cli.QueryRow("SELECT id, name, private_key, create_time, update_time FROM app WHERE id = ?", appId).Scan(
		&app.Id, &app.Name, &app.PrivateKey, &app.CreateTime, &app.UpdateTime)
	if err != nil {
		return nil, err
	}
	return &app, nil
}
