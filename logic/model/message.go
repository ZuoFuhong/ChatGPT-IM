package model

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"strconv"
)

const (
	MessageObjectTypeUser  = 1 // 用户
	MessageObjectTypeGroup = 2 // 群组
)

// Message 消息
type Message struct {
	ObjectType     int    `json:"object_type"`      // 所属类型 1：用户；2：群组
	ObjectId       int64  `json:"object_id"`        // 所属类型ID
	RequestId      int64  `json:"request_id"`       // 请求ID
	SenderType     int32  `json:"sender_type"`      // 发送者类型
	SenderId       int64  `json:"sender_id"`        // 发送者账户ID
	SenderDeviceId int64  `json:"sender_device_id"` // 发送者设备ID
	ReceiverType   int32  `json:"receiver_type"`    // 接收者类型 1:个人；2：普通群组；3：超大群组
	ReceiverId     int64  `json:"receiver_id"`      // 接受者ID，如果是单聊信息，则为user_id，如果是群组消息，则为group_id
	ToUserIds      string `json:"to_user_ids"`      // 需要@的用户ID列表，多个用户 , 隔开
	Type           int    `json:"type"`             // 消息类型
	Content        string `json:"content"`          // 消息内容
	Seq            int64  `json:"seq"`              // 消息同步序列
	SendTime       int64  `json:"send_time"`        // 消息发送时间
	Status         int32  `json:"status"`           // 创建时间
}

func bucketName(recType int32, receiverId int64) string {
	return fmt.Sprintf("message_%d_%d", recType, receiverId)
}

// StoreMessage 存储消息
func StoreMessage(msg *Message) error {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	// 用户维度分片
	bkName := bucketName(msg.ReceiverType, msg.ObjectId)
	return WriteToDB(bkName, fmt.Sprint(msg.Seq), bytes)
}

// ListBySeq 查询消息
func ListBySeq(recType int32, uid, seq int64) ([]*Message, error) {
	bkName := bucketName(recType, uid)
	// 检查 bucket
	if err := getDb().Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bkName))
		return err
	}); err != nil {
		return nil, err
	}
	// 扫描 bucket
	msgList := make([]*Message, 0)
	if err := getDb().View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bkName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			kseq, _ := strconv.ParseInt(string(k), 10, 64)
			if seq >= kseq {
				continue
			}
			msg := &Message{}
			_ = json.Unmarshal(v, msg)
			msgList = append(msgList, msg)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return msgList, nil
}
