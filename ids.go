package offline

import (
	"context"
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {

	n, err := snowflake.NewNode(1)

	if err != nil {
		panic(err)
	}

	node = n
}

// NewJobId returns a unique identifier to be associated with a new job.
func NewJobId(ctx context.Context) (int64, error) {
	id := node.Generate()
	return id.Int64(), nil
}
