package http_handler

import (
	"ChatGPT-IM/backend/consts"
	model2 "ChatGPT-IM/backend/logic/model"
	service2 "ChatGPT-IM/backend/logic/service"
	defs2 "ChatGPT-IM/backend/pkg/defs"
	"ChatGPT-IM/backend/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type user struct{}

var User = new(user)

// Register 注册用户
func (*user) Register(w http.ResponseWriter, r *http.Request) {
	form := new(defs2.RegisterUserForm)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		defs2.Error(w, defs2.ParameterError, "参数错误")
		return
	}
	// 注册用户
	um := &model2.User{
		Nickname:  form.Nickname,
		Sex:       form.Sex,
		AvatarUrl: form.AvatarUrl,
		Extra:     form.Extra,
	}
	userId, err := service2.UserService.Add(um)
	if err != nil {
		defs2.Error(w, defs2.User, "用户注册失败")
		return
	}
	// 注册设备
	dm := new(model2.Device)
	dm.UserId = userId
	dm.Type = model2.Web
	deviceId, err := service2.DeviceService.Register(dm)
	if err != nil {
		defs2.Error(w, defs2.User, "设备注册失败")
		return
	}
	resp := make(map[string]string)
	resp["user_id"] = fmt.Sprint(userId)
	resp["device_id"] = fmt.Sprint(deviceId)
	defs2.Ok(w, resp)
}

// Info 查询用户
func (*user) Info(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["uid"], 10, 64)
	um, err := service2.UserService.UserInfo(userId)
	if err != nil {
		defs2.Error(w, defs2.User, "查询用户信息失败")
		return
	}
	if um == nil {
		defs2.Error(w, defs2.User, "The user does not exist")
		return
	}
	userVO := new(defs2.UserVO)
	userVO.UserId = strconv.FormatInt(um.Id, 10)
	userVO.Nickname = um.Nickname
	userVO.AvatarUrl = um.AvatarUrl
	userVO.Extra = um.Extra

	// 查询 Web 设备
	dm, err := service2.DeviceService.GetUserWebDevice(userId)
	if err != nil {
		defs2.Error(w, defs2.User, "查询用户设备失败")
		return
	}
	token, _ := util.GetToken(userId, dm.Id, time.Now().Add(1*time.Hour).Unix(), consts.PublicKey)

	resp := make(map[string]interface{})
	resp["user"] = userVO
	resp["deviceId"] = fmt.Sprint(dm.Id)
	resp["token"] = token
	defs2.Ok(w, resp)
}

// Update 更新用户
func (*user) Update(w http.ResponseWriter, r *http.Request) {
	form := new(defs2.UpdateUserForm)
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		defs2.Error(w, defs2.ParameterError, "参数错误")
		return
	}
	um := new(model2.User)
	um.Id, _ = strconv.ParseInt(form.UserId, 10, 64)
	um.Nickname = form.Nickname
	um.Sex = form.Sex
	um.AvatarUrl = form.AvatarUrl
	um.Extra = form.Extra
	if err := service2.UserService.Update(um); err != nil {
		defs2.Error(w, defs2.User, "更新用户信息失败")
		return
	}
	defs2.Ok(w, "ok")
}
