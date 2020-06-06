package http_handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-IM/logic/service"
	"go-IM/pkg/defs"
	"net/http"
	"strconv"
)

type friend struct{}

var Friend = new(friend)

// 添加好友
func (*friend) AddFriend(w http.ResponseWriter, r *http.Request) {
	friend := new(defs.AddFriendForm)
	err := json.NewDecoder(r.Body).Decode(&friend)
	if err != nil {
		panic(err)
	}
	appId, _ := strconv.ParseInt(friend.AppId, 10, 64)
	userId, _ := strconv.ParseInt(friend.UserId, 10, 64)
	friendId, _ := strconv.ParseInt(friend.FriendId, 10, 64)
	service.FriendService.AddFriend(appId, userId, friendId)
	_, _ = w.Write([]byte("ok"))
}

// 查询好友
func (*friend) ListFriend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["uid"], 10, 64)

	friends := service.FriendService.QueryFriends(userId)
	bytes, err := json.Marshal(friends)
	if err != nil {
		panic(err)
	}
	userVOS := make([]defs.UserVO, 0, 5)
	for _, v := range *friends {
		userFriend := new(defs.UserVO)
		userFriend.UserId = strconv.FormatInt(v.UserId, 10)
		userFriend.Nickname = v.Nickname
		userFriend.AvatarUrl = v.AvatarUrl
		userFriend.Extra = v.Extra
		userVOS = append(userVOS, *userFriend)
	}
	bytes, err = json.Marshal(userVOS)
	if err != nil {
		panic(err)
	}
	_, _ = w.Write(bytes)
}
