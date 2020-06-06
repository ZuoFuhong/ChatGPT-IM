package defs

type PackageType int

const (
	PackageType_SignIn      PackageType = 1
	PackageType_SYNC        PackageType = 2
	PackageType_HEARTBEAT   PackageType = 3
	PackageType_MESSAGE_ACK PackageType = 4
	PackageType_RT_USER     PackageType = 5
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
	AppId    string
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

// 发送者类型
type SenderType int32

const (
	SenderType_ST_SYSTEM SenderType = 1 // 系统
	SenderType_ST_USER   SenderType = 2 // 用户
)

type Sender struct {
	AppId      int64      // appId
	SenderType SenderType // 发送者类型
	SenderId   int64      // 发送者id
	DeviceId   int64      // 发送者设备id
}

type MessageType int32

const (
	Message_TEXT  MessageType = 1
	Message_IMAGE MessageType = 2
	Message_AUDIO MessageType = 3
)

// 接收者类型
type ReceiverType int32

const (
	ReceiverType_RT_USER         ReceiverType = 1 // 用户
	ReceiverType_RT_NORMAL_GROUP ReceiverType = 2 // 普通群组
	ReceiverType_RT_LARGE_GROUP  ReceiverType = 3 // 超大群组
)

type SendMessageReq struct {
	ReceiverType   ReceiverType
	ReceiverId     int64
	MessageType    MessageType
	MessageContent string
	ToUserIds      []string
	IsPersist      bool
}

// 消息状态
type MessageStatus int32

const (
	MessageStatus_MS_NORMAL MessageStatus = 1 // 未处理
	MessageStatus_MS_RECALL MessageStatus = 2 // 消息撤回
)

type SendMessage struct {
	AppId          string       // appId
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
