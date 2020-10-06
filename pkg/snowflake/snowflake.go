package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

// 雪花算法，生成id值

var node *snowflake.Node

// 开始时间 + 机器Id
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
