package dynamodb

import (
	"context"
	"fmt"
	"github.com/aaronland/go-aws-session"
	"github.com/aws/aws-sdk-go/aws"
	aws_dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	"net/url"
	"os"
	"strconv"
)

func NewClientWithURI(ctx context.Context, uri string) (*aws_dynamodb.DynamoDB, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()
	region := q.Get("region")
	credentials := q.Get("credentials")
	local := q.Get("local")

	is_local, err := strconv.ParseBool(local)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse ?local parameter, %w", err)
	}

	if is_local {
		os.Setenv("AWS_ACCESS_KEY_ID", "DUMMYIDEXAMPLE")
		os.Setenv("AWS_SECRET_ACCESS_KEY", "DUMMYEXAMPLEKEY")
		credentials = "env:"
		region = "us-east-1"
	}

	dsn := fmt.Sprintf("credentials=%s region=%s", credentials, region)

	sess, err := session.NewSessionWithDSN(dsn)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new session, %w", err)
	}

	if is_local {
		endpoint := "http://localhost:8000"
		sess.Config.Endpoint = aws.String(endpoint)
	}

	client := aws_dynamodb.New(sess)
	return client, nil
}
