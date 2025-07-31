package offline

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	_ "gocloud.dev/pubsub/awssnssqs"
	_ "gocloud.dev/pubsub/mempubsub"

	"github.com/sfomuseum/go-pubsub/publisher"
)

type PubSubQueue struct {
	Queue
	publisher publisher.Publisher
}

func init() {

	ctx := context.Background()

	err := publisher.RegisterGoCloudPublishers(ctx)

	if err != nil {
		panic(err)
	}

	for _, uri := range publisher.PublisherSchemes() {

		scheme := strings.Replace(uri, "://", "", 1)

		// Skip go-pubsub/null_publisher in favour of go-offline/null-queue

		if scheme == "null" {
			continue
		}

		err := RegisterQueue(ctx, scheme, NewPubSubQueue)

		if err != nil {
			panic(err)
		}
	}
}

func NewPubSubQueue(ctx context.Context, uri string) (Queue, error) {

	p, err := publisher.NewPublisher(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create publisher, %w", err)
	}

	q := &PubSubQueue{
		publisher: p,
	}

	return q, nil
}

func (q *PubSubQueue) ScheduleJob(ctx context.Context, job_id int64) error {

	str_id := strconv.FormatInt(job_id, 10)

	err := q.publisher.Publish(ctx, str_id)

	if err != nil {
		return fmt.Errorf("Failed to send message, %w", err)
	}

	return nil
}

func (q *PubSubQueue) Close(ctx context.Context) error {

	err := q.publisher.Close()

	if err != nil {
		return fmt.Errorf("Failed to close publisher, %w", err)
	}

	return nil
}
