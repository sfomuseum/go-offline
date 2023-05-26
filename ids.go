package offline

import (
	"context"
	"math/rand"
	
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {

	node_id := rand.Int63n(1023)
	n, err := snowflake.NewNode(node_id)

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
