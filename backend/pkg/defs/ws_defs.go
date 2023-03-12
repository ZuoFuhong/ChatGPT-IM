package defs

type PackageType int

const (
	PackagetypeSignin     PackageType = 1
	PackagetypeSync       PackageType = 2
	PackagetypeHeartbeat  PackageType = 3
	PackagetypeMessageAck PackageType = 4
	PackagetypeRtUser     PackageType = 5
)

type Input struct {
	Type      PackageType
	RequestId int64
	Data      string
}

type Output struct {
	Type      PackageType
	RequestId int64
	Code      int
	Message   string
	Data      interface{}
}

type SignIn struct {
	UserId   string
	DeviceId string
	Token    string
}

type SyncInput struct {
	Seq string
}

type SyncOutput struct {
	Messages []MessageItem
}

type MessageACK struct {
	DeviceAck   int64
	ReceiveTime int64
}

// SenderType 发送者类型
type SenderType int32

const (
	SendertypeStSystem SenderType = 1 // 系统
	SendertypeStUser   SenderType = 2 // 用户
)

type Sender struct {
	SenderType SenderType // 发送者类型
	SenderId   int64      // 发送者ID
	DeviceId   int64      // 发送者设备ID
}

type MessageType int32

const (
	MessageText  MessageType = 1
	MessageImage MessageType = 2
	MessageAudio MessageType = 3
)

// ReceiverType 接收者类型
type ReceiverType int32

const (
	ReceivertypeRtUser        ReceiverType = 1 // 用户
	ReceivertypeRtNormalGroup ReceiverType = 2 // 普通群组
	ReceivertypeRtLargeGroup  ReceiverType = 3 // 超大群组
)

type SendMessageReq struct {
	ReceiverType   ReceiverType
	ReceiverId     int64
	MessageType    MessageType
	MessageContent string
	ToUserIds      []string
	IsPersist      bool
}

// MessageStatus 消息状态
type MessageStatus int32

const (
	MessagestatusMsNormal MessageStatus = 1 // 未处理
	MessageStatusMsRecall MessageStatus = 2 // 消息撤回
)

type SendMessage struct {
	SenderId       string       // 发送者id
	DeviceId       string       // 发送者设备id
	ReceiverType   ReceiverType // 接收者类型
	ReceiverId     string       // 接收者ID
	MessageType    MessageType  // 消息类型
	MessageContent string       // 消息内容
	ToUserIds      []string     // 需要@的用户
}

type MessageItem struct {
	SenderId         string      // 发送者id
	ReceiverId       string      // 接收者id
	ReceiverDeviceId int64       // 接收者设备ID
	SendTime         string      // 发送时间
	Type             MessageType // 消息类型
	Content          string      // 消息内容
	Seq              string      // 序列号
}
