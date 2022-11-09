package offline

import (
	"context"
	"fmt"
	"gocloud.dev/pubsub"
	"strconv"
)

type PubSubQueue struct {
	Queue
	topic *pubsub.Topic
}

func init() {

	ctx := context.Background()

	for _, scheme := range pubsub.DefaultURLMux().TopicSchemes() {
		RegisterQueue(ctx, scheme, NewPubSubQueue)
	}
}

func NewPubSubQueue(ctx context.Context, uri string) (Queue, error) {

	t, err := pubsub.OpenTopic(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to open topic, %w", err)
	}

	q := &PubSubQueue{
		topic: t,
	}

	return q, nil
}

func (q *PubSubQueue) QueueJob(ctx context.Context, job_id int64) error {

	str_id := strconv.FormatInt(job_id, 10)

	msg := &pubsub.Message{
		Body:     []byte(str_id),
		Metadata: map[string]string{
			// "language":   "en",
		},
	}

	err := q.topic.Send(ctx, msg)

	if err != nil {
		return fmt.Errorf("Failed to send message, %w", err)
	}

	return nil
}

func (q *PubSubQueue) Close(ctx context.Context) error {

	err := q.topic.Shutdown(ctx)

	if err != nil {
		return fmt.Errorf("Failed to shutdown topic, %w", err)
	}

	return nil
}
