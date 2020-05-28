package web

import (
	"log"
	"net/http"
)

type App struct {
	router *router
}

func NewApp() *App {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := new(App)
	app.router = NewRouter()
	return app
}

func (app *App) Run(addr string) {
	log.Print("server runs on http://" + addr)
	app.router.RegisterHandler()
	err := http.ListenAndServe(addr, app.router.muxr)
	if err != nil {
		log.Panic(err)
	}
}
