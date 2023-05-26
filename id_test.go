package offline

import (
	"context"
	"testing"
)

func TestNewJobId(t *testing.T) {

	SNOWFLAKE_NODE_ID = int64(674)

	ctx := context.Background()
	_, err := NewJobId(ctx)

	if err != nil {
		t.Fatalf("Failed to create new ID, %v", err)
	}
}
