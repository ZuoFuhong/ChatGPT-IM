package tinyid

import (
	"github.com/bwmarrin/snowflake"
	"log"
	"sync"
)

var nodeIns *snowflake.Node
var once sync.Once

func getNodeIns() *snowflake.Node {
	once.Do(func() {
		node, err := snowflake.NewNode(1)
		if err != nil {
			log.Printf("snowflake.NewNode failed, err: %v\n", err)
		}
		nodeIns = node
	})
	return nodeIns
}

// NextId 生成唯一ID
func NextId() int64 {
	return getNodeIns().Generate().Int64()
}
