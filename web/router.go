package web

import (
	"github.com/gorilla/mux"
	"go-IM/handler"
	"go-IM/internal/ws_conn"
	"go-IM/pkg/alice"
)

type router struct {
	muxr *mux.Router
}

func NewRouter() *router {
	var router = new(router)
	router.muxr = mux.NewRouter()
	return router
}

func (r *router) RegisterHandler() {
	mw := &Middleware{}
	chain := alice.New(mw.LoggingHandler, mw.RecoverPanic, mw.CORSHandler)
	r.muxr.Handle("/ws", chain.ThenFunc(ws_conn.WSHandler)).Methods("GET")
	r.muxr.Handle("/device/register", chain.ThenFunc(handler.Device.Register)).Methods("POST")
	r.muxr.Handle("/user/register", chain.ThenFunc(handler.User.Register)).Methods("POST")
	r.muxr.Handle("/user/info", chain.ThenFunc(handler.User.Info)).Methods("GET").Queries("aid", "{aid}").Queries("uid", "{uid}")
	r.muxr.Handle("/user/update", chain.ThenFunc(handler.User.Update)).Methods("PUT")
	r.muxr.Handle("/group/create", chain.ThenFunc(handler.Group.Create)).Methods("POST")
	r.muxr.Handle("/group/info", chain.ThenFunc(handler.Group.Info)).Methods("GET").Queries("aid", "{aid}").Queries("uid", "{uid}")
	r.muxr.Handle("/group/update", chain.ThenFunc(handler.Group.Update)).Methods("PUT")
	r.muxr.Handle("/group/users", chain.ThenFunc(handler.Group.Users)).Methods("GET").Queries("aid", "{aid}").Queries("gid", "{gid}")
	r.muxr.Handle("/group/user/add", chain.ThenFunc(handler.Group.AddUser)).Methods("POST")
	r.muxr.Handle("/group/user/update", chain.ThenFunc(handler.Group.Update)).Methods("PUT")
	r.muxr.Handle("/group/user/groups", chain.ThenFunc(handler.Group.Groups)).Methods("GET").Queries("aid", "{aid}").Queries("uid", "{uid}")
	r.muxr.Handle("/group/user/kick", chain.ThenFunc(handler.Group.KickUser)).Methods("PUT").Queries("aid", "{aid}").Queries("gid", "{gid}").Queries("uid", "{uid}")
}
