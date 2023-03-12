package web

import (
	"log"
	"net/http"
)

type App struct {
	router *Router
}

func NewApp() *App {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := new(App)
	app.router = NewRouter()
	return app
}

func (app *App) Run(address string) {
	log.Print("server runs on http://" + address)
	app.router.RegisterHandler()
	err := http.ListenAndServe(address, app.router.muxr)
	if err != nil {
		log.Panic(err)
	}
}
