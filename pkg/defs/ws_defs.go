package defs

import "go-IM/internal/logic/model"

type PackageType int

const (
	PackageType_SYNC      PackageType = 1
	PackageType_HEARTBEAT PackageType = 2
	PackageType_MESSAGE   PackageType = 3
)

type Input struct {
	Type      PackageType
	RequestId int
	Data      []byte
}

type Output struct {
	Type      PackageType
	RequestId int
	Code      int
	Message   string
	Data      []byte
}

type SyncInput struct {
	Seq int64
}

type SyncOutput struct {
	Messages []model.Message
}

type MessageACK struct {
	DeviceAck   int64
	ReceiveTime int64
}
