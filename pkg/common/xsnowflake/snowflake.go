package xsnowflake

import (
	"github.com/bwmarrin/snowflake"
	"lark/pkg/utils"
	"math/rand"
	"strconv"
)

var (
	Snowflake     *snowflakeNode
	maxNodeNumber = 1023
)

type snowflakeNode struct {
	node *snowflake.Node
}

func NewSnowflake(n int) {
	var (
		node *snowflake.Node
		err  error
	)
	if n < 0 || n > maxNodeNumber {
		n = 1
	}
	// Create a new Node with a Node number of 1
	node, err = snowflake.NewNode(int64(n))
	if err != nil {
		return
	}
	Snowflake = &snowflakeNode{node}
}

// Generate a snowflake ID.
func NewSnowflakeID() int64 {
	if Snowflake == nil {
		NewSnowflake(rand.Intn(maxNodeNumber))
	}
	return Snowflake.node.Generate().Int64()
}

func NewStrSnowflakeID() string {
	if Snowflake == nil {
		NewSnowflake(rand.Intn(maxNodeNumber))
	}
	return Snowflake.node.Generate().String()
}

func DefaultLarkId() string {
	if Snowflake == nil {
		NewSnowflake(rand.Intn(maxNodeNumber))
	}
	return utils.SixteenMD5(strconv.FormatInt(NewSnowflakeID(), 10))
}
