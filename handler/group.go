package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-IM/internal/logic/model"
	"go-IM/internal/logic/service"
	"go-IM/pkg/defs"
	"go-IM/pkg/errs"
	"net/http"
	"strconv"
)

type group struct{}

var Group = new(group)

// 创建群组
func (*group) Create(w http.ResponseWriter, r *http.Request) {
	form := new(defs.CreateGroupForm)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		panic(errs.ParameterError)
	}
	group := new(model.Group)
	group.AppId = form.AppId
	group.Name = form.Name
	group.Introduction = form.Introduction
	group.UserNum = form.UserNum
	group.Type = form.Type
	group.Extra = form.Extra
	groupId := service.GroupService.Create(group)

	resp := make(map[string]interface{})
	resp["groupId"] = groupId
	bytes, _ := json.Marshal(resp)
	_, _ = w.Write(bytes)
}

// 群组信息
func (*group) Info(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appId, _ := strconv.ParseInt(vars["aid"], 10, 64)
	groupId, _ := strconv.ParseInt(vars["gid"], 10, 64)
	group := service.GroupService.Get(appId, groupId)

	bytes, _ := json.Marshal(group)
	_, _ = w.Write(bytes)
}

// 更新群组
func (*group) Update(w http.ResponseWriter, r *http.Request) {
	form := new(defs.UpdateGroupForm)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		panic(errs.ParameterError)
	}

	group := new(model.Group)
	group.AppId = form.AppId
	group.GroupId = form.GroupId
	group.Name = form.Name
	group.Introduction = form.Introduction
	group.Extra = form.Extra
	service.GroupService.Update(group)

	_, _ = w.Write([]byte("ok"))
}

// 查询群组用户
func (*group) Users(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appId, _ := strconv.ParseInt(vars["aid"], 10, 64)
	groupId, _ := strconv.ParseInt(vars["gid"], 10, 64)
	users := service.GroupUserService.GetUsers(appId, groupId)

	bytes, _ := json.Marshal(users)
	_, _ = w.Write(bytes)
}

// 群组添加用户
func (*group) AddUser(w http.ResponseWriter, r *http.Request) {
	form := new(defs.GroupAddUserForm)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		panic(errs.ParameterError)
	}
	groupUser := new(model.GroupUser)
	groupUser.AppId = form.AppId
	groupUser.GroupId = form.GroupId
	groupUser.UserId = form.UserId
	groupUser.Label = form.Label
	groupUser.Extra = form.Extra
	service.GroupService.AddUser(groupUser)

	_, _ = w.Write([]byte("ok"))
}

// 更新群组用户
func (*group) UpdateUser(w http.ResponseWriter, r *http.Request) {
	form := new(defs.UpdateGroupUserForm)
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		panic(errs.ParameterError)
	}
	groupUser := new(model.GroupUser)
	groupUser.AppId = form.AppId
	groupUser.GroupId = form.GroupId
	groupUser.UserId = form.UserId
	groupUser.Label = form.Label
	groupUser.Extra = form.Extra
	service.GroupService.UpdateUser(groupUser)

	_, _ = w.Write([]byte("ok"))
}

// 查询用户加入的群组
func (*group) Groups(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appId, _ := strconv.ParseInt(vars["aid"], 10, 64)
	userId, _ := strconv.ParseInt(vars["uid"], 10, 64)
	groups := service.GroupUserService.ListByUserId(appId, userId)

	bytes, _ := json.Marshal(groups)
	_, _ = w.Write(bytes)
}

// 群组踢人
func (*group) KickUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appId, _ := strconv.ParseInt(vars["aid"], 10, 64)
	groupId, _ := strconv.ParseInt(vars["gid"], 10, 64)
	userId, _ := strconv.ParseInt(vars["uid"], 10, 64)

	service.GroupService.DeleteUser(appId, groupId, userId)
	_, _ = w.Write([]byte("ok"))
}
