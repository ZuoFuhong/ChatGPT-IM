package web

import (
	"github.com/gorilla/mux"
	"go-IM/pkg/alice"
	"go-IM/web/http_handler"
	"go-IM/web/ws_handler"
	"net/http"
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
	r.muxr.Handle("/device/register", chain.ThenFunc(http_handler.Device.Register)).Methods("POST")
	r.muxr.Handle("/user/register", chain.ThenFunc(http_handler.User.Register)).Methods("POST")
	r.muxr.Handle("/user/info", chain.ThenFunc(http_handler.User.Info)).Methods("GET").Queries("aid", "{aid}").Queries("uid", "{uid}")
	r.muxr.Handle("/user/update", chain.ThenFunc(http_handler.User.Update)).Methods("PUT")
	r.muxr.Handle("/group/create", chain.ThenFunc(http_handler.Group.Create)).Methods("POST")
	r.muxr.Handle("/group/info", chain.ThenFunc(http_handler.Group.Info)).Methods("GET").Queries("aid", "{aid}").Queries("gid", "{gid}")
	r.muxr.Handle("/group/update", chain.ThenFunc(http_handler.Group.Update)).Methods("PUT")
	r.muxr.Handle("/group/users", chain.ThenFunc(http_handler.Group.Users)).Methods("GET").Queries("aid", "{aid}").Queries("gid", "{gid}")
	r.muxr.Handle("/group/user/add", chain.ThenFunc(http_handler.Group.AddUser)).Methods("POST")
	r.muxr.Handle("/group/user/update", chain.ThenFunc(http_handler.Group.UpdateUser)).Methods("PUT")
	r.muxr.Handle("/group/user/groups", chain.ThenFunc(http_handler.Group.Groups)).Methods("GET").Queries("aid", "{aid}").Queries("uid", "{uid}")
	r.muxr.Handle("/group/user/kick", chain.ThenFunc(http_handler.Group.KickUser)).Methods("DELETE").Queries("aid", "{aid}").Queries("gid", "{gid}").Queries("uid", "{uid}")
	r.muxr.Handle("/ws", chain.ThenFunc(ws_handler.WSHandler)).Methods("GET")
	// 静态资源
	r.muxr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./"))))
}
