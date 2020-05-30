package web

import (
	"go-IM/config"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	router *router
	conf   *config.Conf
}

func NewApp() *App {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := new(App)
	app.router = NewRouter()
	app.conf = config.LoadConf()
	return app
}

func (app *App) Run() {
	addr := app.conf.Server.Addr + ":" + strconv.Itoa(app.conf.Server.Port)
	log.Print("server runs on http://" + addr)
	app.router.RegisterHandler()
	err := http.ListenAndServe(addr, app.router.muxr)
	if err != nil {
		log.Panic(err)
	}
}
