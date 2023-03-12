package main

import (
	"ChatGPT-IM/backend/consts"
	"ChatGPT-IM/backend/web"
)

func main() {
	web.NewApp().Run(consts.DefaultAddress)
}
