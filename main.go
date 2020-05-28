package main

import "go-IM/web"

func main() {
	web.NewApp().Run("127.0.0.1:8080")
}
