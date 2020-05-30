package http_handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-IM/logic/model"
	"go-IM/logic/service"
	"go-IM/pkg/defs"
	"go-IM/pkg/errs"
	"net/http"
	"strconv"
)

type user struct{}

var User = new(user)

// 注册用户
func (*user) Register(w http.ResponseWriter, r *http.Request) {
	form := new(defs.RegisterUserForm)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		panic(errs.ParameterError)
	}
	user := new(model.User)
	user.AppId = form.AppId
	user.Nickname = form.Nickname
	user.Sex = form.Sex
	user.AvatarUrl = form.AvatarUrl
	user.Extra = form.Extra
	userId := service.UserService.Add(user)

	resp := make(map[string]interface{})
	resp["userId"] = userId
	bytes, _ := json.Marshal(resp)
	_, _ = w.Write(bytes)
}

// 查询用户
func (*user) Info(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appId, _ := strconv.ParseInt(vars["aid"], 10, 64)
	userId, _ := strconv.ParseInt(vars["uid"], 10, 64)

	user := service.UserService.Get(appId, userId)
	bytes, _ := json.Marshal(user)
	_, _ = w.Write(bytes)
}

// 更新用户
func (*user) Update(w http.ResponseWriter, r *http.Request) {
	form := new(defs.UpdateUserForm)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		panic(errs.ParameterError)
	}
	user := new(model.User)
	user.AppId = form.AppId
	user.UserId = form.UserId
	user.Nickname = form.Nickname
	user.Sex = form.Sex
	user.AvatarUrl = form.AvatarUrl
	user.Extra = form.Extra
	service.UserService.Update(user)

	_, _ = w.Write([]byte("ok"))
}
