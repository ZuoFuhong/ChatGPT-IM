package ws_conn

import "sync"

var manager sync.Map

func store(deviceId int64, ctx *ConnContext) {
	manager.Store(deviceId, ctx)
}

func load(deviceId int64) *ConnContext {
	value, ok := manager.Load(deviceId)
	if ok {
		return value.(*ConnContext)
	}
	return nil
}

func delete(deviceId int64) {
	manager.Delete(deviceId)
}
