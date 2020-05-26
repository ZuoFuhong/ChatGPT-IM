package service

import (
	"go-IM/internal/logic/dao"
	"go-IM/internal/logic/model"
)

type appService struct{}

var AppService = new(appService)

func (*appService) Get(appId int64) (*model.App, error) {
	return dao.AppDao.Get(appId)
}
