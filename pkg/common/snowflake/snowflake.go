package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

var Snowflake *snowflakeNode

type snowflakeNode struct {
	node *snowflake.Node
}

func init() {
	var (
		node *snowflake.Node
		err  error
	)
	// Create a new Node with a Node number of 1
	node, err = snowflake.NewNode(1)
	if err != nil {
		return
	}
	Snowflake = &snowflakeNode{node}
}

// Generate a snowflake ID.
func SnowflakeID() (id string) {
	id = Snowflake.node.Generate().String()
	return
}
