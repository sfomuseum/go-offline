package offline

import (
	"context"
	"testing"
)

func TestNewId(t *testing.T) {

	ctx := context.Background()
	_, err := NewJobId(ctx)

	if err != nil {
		t.Fatalf("Failed to create new ID, %v", err)
	}
}
