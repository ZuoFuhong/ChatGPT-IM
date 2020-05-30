package service

import (
	"go-IM/logic/dao"
	"go-IM/logic/model"
)

type appService struct{}

var AppService = new(appService)

func (*appService) Get(appId int64) (*model.App, error) {
	return dao.AppDao.Get(appId)
}
