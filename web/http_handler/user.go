package http_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-IM/consts"
	"go-IM/logic/model"
	"go-IM/logic/service"
	"go-IM/pkg/defs"
	"go-IM/pkg/util"
	"net/http"
	"strconv"
	"time"
)

type user struct{}

var User = new(user)

// Register 注册用户
func (*user) Register(w http.ResponseWriter, r *http.Request) {
	form := new(defs.RegisterUserForm)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		defs.Error(w, defs.ParameterError, "参数错误")
		return
	}
	// 注册用户
	um := &model.User{
		Nickname:  form.Nickname,
		Sex:       form.Sex,
		AvatarUrl: form.AvatarUrl,
		Extra:     form.Extra,
	}
	userId, err := service.UserService.Add(um)
	if err != nil {
		defs.Error(w, defs.User, "用户注册失败")
		return
	}
	// 注册设备
	dm := new(model.Device)
	dm.UserId = userId
	dm.Type = model.Web
	deviceId, err := service.DeviceService.Register(dm)
	if err != nil {
		defs.Error(w, defs.User, "设备注册失败")
		return
	}
	resp := make(map[string]string)
	resp["user_id"] = fmt.Sprint(userId)
	resp["device_id"] = fmt.Sprint(deviceId)
	defs.Ok(w, resp)
}

// Info 查询用户
func (*user) Info(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["uid"], 10, 64)
	um, err := service.UserService.UserInfo(userId)
	if err != nil {
		defs.Error(w, defs.User, "查询用户信息失败")
		return
	}
	if um == nil {
		defs.Error(w, defs.User, "The user does not exist")
		return
	}
	userVO := new(defs.UserVO)
	userVO.UserId = strconv.FormatInt(um.Id, 10)
	userVO.Nickname = um.Nickname
	userVO.AvatarUrl = um.AvatarUrl
	userVO.Extra = um.Extra

	// 查询 Web 设备
	dm, err := service.DeviceService.GetUserWebDevice(userId)
	if err != nil {
		defs.Error(w, defs.User, "查询用户设备失败")
		return
	}
	token, _ := util.GetToken(userId, dm.Id, time.Now().Add(1*time.Hour).Unix(), consts.PublicKey)

	resp := make(map[string]interface{})
	resp["user"] = userVO
	resp["deviceId"] = fmt.Sprint(dm.Id)
	resp["token"] = token
	defs.Ok(w, resp)
}

// Update 更新用户
func (*user) Update(w http.ResponseWriter, r *http.Request) {
	form := new(defs.UpdateUserForm)
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		defs.Error(w, defs.ParameterError, "参数错误")
		return
	}
	um := new(model.User)
	um.Id, _ = strconv.ParseInt(form.UserId, 10, 64)
	um.Nickname = form.Nickname
	um.Sex = form.Sex
	um.AvatarUrl = form.AvatarUrl
	um.Extra = form.Extra
	if err := service.UserService.Update(um); err != nil {
		defs.Error(w, defs.User, "更新用户信息失败")
		return
	}
	defs.Ok(w, "ok")
}
