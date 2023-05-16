package offline

import (
	"context"
	"net/url"

	"github.com/aaronland/go-roster"
)

type Queue interface {
	ScheduleJob(context.Context, int64) error
	Close(context.Context) error
}

var queue_roster roster.Roster

// QueueInitializationFunc is a function defined by individual queue package and used to create
// an instance of that queue
type QueueInitializationFunc func(ctx context.Context, uri string) (Queue, error)

// RegisterQueue registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `Queue` instances by the `NewQueue` method.
func RegisterQueue(ctx context.Context, scheme string, init_func QueueInitializationFunc) error {

	err := ensureQueueRoster()

	if err != nil {
		return err
	}

	return queue_roster.Register(ctx, scheme, init_func)
}

func ensureQueueRoster() error {

	if queue_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		queue_roster = r
	}

	return nil
}

// NewQueue returns a new `Queue` instance configured by 'uri'. The value of 'uri' is parsed
// as a `url.URL` and its scheme is used as the key for a corresponding `QueueInitializationFunc`
// function used to instantiate the new `Queue`. It is assumed that the scheme (and initialization
// function) have been registered by the `RegisterQueue` method.
func NewQueue(ctx context.Context, uri string) (Queue, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := queue_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(QueueInitializationFunc)
	return init_func(ctx, uri)
}

/*
// Schemes returns the list of schemes that have been registered.
func Schemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureQueueRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range queue_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}

*/
