package http_handler

import (
	"ChatGPT-IM/backend/logic/service"
	defs2 "ChatGPT-IM/backend/pkg/defs"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type friend struct{}

var Friend = new(friend)

// AddFriend 添加好友
func (*friend) AddFriend(w http.ResponseWriter, r *http.Request) {
	form := new(defs2.AddFriendForm)
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		defs2.Error(w, defs2.ParameterError, "参数错误")
		return
	}
	userId, _ := strconv.ParseInt(form.UserId, 10, 64)
	friendId, _ := strconv.ParseInt(form.FriendId, 10, 64)
	if err := service.FriendService.AddFriend(userId, friendId); err != nil {
		defs2.Error(w, defs2.Friend, "添加好友失败")
		return
	}
	defs2.Ok(w, "ok")
}

// ListFriend 查询好友
func (*friend) ListFriend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["uid"], 10, 64)

	friends, err := service.FriendService.QueryFriends(userId)
	if err != nil {
		defs2.Error(w, defs2.Friend, "查询好友失败")
		return
	}
	vmvoList := make([]*defs2.UserVO, 0)
	for _, um := range friends {
		umvo := &defs2.UserVO{
			UserId:    strconv.FormatInt(um.Id, 10),
			Nickname:  um.Nickname,
			AvatarUrl: um.AvatarUrl,
			Extra:     um.Extra,
		}
		vmvoList = append(vmvoList, umvo)
	}
	defs2.Ok(w, vmvoList)
}
