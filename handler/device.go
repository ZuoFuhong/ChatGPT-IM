package handler

import (
	"encoding/json"
	"go-IM/internal/logic/model"
	"go-IM/internal/logic/service"
	"go-IM/pkg/defs"
	"go-IM/pkg/errs"
	"net/http"
)

type device struct{}

var Device = new(device)

// 注册设备
func (*device) Register(w http.ResponseWriter, r *http.Request) {
	var form = new(defs.RegisterDeviceForm)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		panic(errs.ParameterError)
	}
	device := new(model.Device)
	device.AppId = form.AppId
	device.UserId = form.UserId
	device.Type = form.Type
	device.Brand = form.Brand
	device.Model = form.Model
	device.SystemVersion = form.SystemVersion
	device.SDKVersion = form.SDKVersion
	device.Status = form.Status
	deviceId, err := service.DeviceService.Register(device)
	if err != nil {
		panic(errs.NewHttpErr(errs.Device, "注册失败"))
	}
	resp := make(map[string]interface{})
	resp["deviceId"] = deviceId
	bytes, _ := json.Marshal(resp)
	_, _ = w.Write(bytes)
}
