package generate

import (
	"flag"
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	envSnowflakeStartTime = flag.Int64(
		"SNOWFLAKE_START_TIME",
		time.UnixMicro(0).UnixMicro(),
		"time from each start generates snowflake id",
	)
	envSnowflakeNodeID = flag.Int64(
		"SNOWFLAKE_NODE_ID",
		0,
		"unique node id for each service's instance",
	)
)

var (
	node *snowflake.Node
)

func init() {
	snowflake.Epoch = *envSnowflakeStartTime
	var (
		nodeID = *envSnowflakeNodeID
		err    error
	)
	node, err = snowflake.NewNode(nodeID)
	if err != nil {
		panic(fmt.Sprintf(
			"can not to create snowflake for node %d, error: %s",
			nodeID, err.Error(),
		))
	}
}

func UniqueID() int64 {
	return node.Generate().Int64()
}
