package main

import WSConn "go-IM/internal/ws_conn"

func main() {
	WSConn.StartServer("127.0.0.1:8080")
}
