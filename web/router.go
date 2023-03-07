package web

import (
	"github.com/gorilla/mux"
	"go-IM/pkg/alice"
	"go-IM/web/http_handler"
	"go-IM/web/ws_handler"
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
	r.muxr.Handle("/user/register", chain.ThenFunc(http_handler.User.Register)).Methods("POST")
	r.muxr.Handle("/user/info", chain.ThenFunc(http_handler.User.Info)).Methods("GET").Queries("uid", "{uid}")
	r.muxr.Handle("/user/update", chain.ThenFunc(http_handler.User.Update)).Methods("PUT")
	r.muxr.Handle("/friend/add", chain.ThenFunc(http_handler.Friend.AddFriend)).Methods("POST")
	r.muxr.Handle("/friend/list", chain.ThenFunc(http_handler.Friend.ListFriend)).Methods("GET").Queries("uid", "{uid}")
	r.muxr.Handle("/audio/transcriptions", chain.ThenFunc(http_handler.Media.AudioTranscriptions)).Methods("POST")
	r.muxr.Handle("/ws", chain.ThenFunc(ws_handler.WSHandler)).Methods("GET")
}
