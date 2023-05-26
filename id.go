package offline

import (
	"context"
	"sync"

	"github.com/bwmarrin/snowflake"
)

var SNOWFLAKE_NODE_ID = int64(575)

var node *snowflake.Node
var node_err error
var node_once sync.Once

// NewJobId returns a unique identifier to be associated with a publication.
func NewJobId(ctx context.Context) (int64, error) {

	node_once.Do(setupSnowflakeNode)

	if node_err != nil {
		return 0, node_err
	}

	id := node.Generate()
	return id.Int64(), nil
}

func setupSnowflakeNode() {
	node, node_err = snowflake.NewNode(SNOWFLAKE_NODE_ID)
}
