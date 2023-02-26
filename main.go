package main

import (
	"go-IM/consts"
	"go-IM/web"
)

func main() {
	web.NewApp().Run(consts.DefaultAddress)
}
