package web

import (
	"ChatGPT-IM/backend/pkg/alice"
	http_handler2 "ChatGPT-IM/backend/web/http_handler"
	"ChatGPT-IM/backend/web/ws_handler"
	"github.com/gorilla/mux"
)

type Router struct {
	muxr *mux.Router
}

func NewRouter() *Router {
	var router = new(Router)
	router.muxr = mux.NewRouter()
	return router
}

func (r *Router) RegisterHandler() {
	mw := NewMiddleware()
	chain := alice.New(mw.CORSHandler)
	r.muxr.Handle("/user/register", chain.ThenFunc(http_handler2.User.Register)).Methods("POST")
	r.muxr.Handle("/user/info", chain.ThenFunc(http_handler2.User.Info)).Methods("GET").Queries("uid", "{uid}")
	r.muxr.Handle("/user/update", chain.ThenFunc(http_handler2.User.Update)).Methods("PUT")
	r.muxr.Handle("/friend/add", chain.ThenFunc(http_handler2.Friend.AddFriend)).Methods("POST")
	r.muxr.Handle("/friend/list", chain.ThenFunc(http_handler2.Friend.ListFriend)).Methods("GET").Queries("uid", "{uid}")
	r.muxr.Handle("/audio/transcriptions", chain.ThenFunc(http_handler2.Media.AudioTranscriptions)).Methods("POST")
	r.muxr.Handle("/ws", chain.ThenFunc(ws_handler.WSHandler)).Methods("GET")
}
