package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoDBTables is a map whose keys are DynamoDB table names and whose values are `dynamodb.CreateTableInput` instances.
var DynamoDBTables = map[string]*dynamodb.CreateTableInput{
	"offlinejobs": &dynamodb.CreateTableInput{
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       "HASH", // partition key
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: "N",
			},
			{
				AttributeName: aws.String("LastModified"),
				AttributeType: "N",
			},
			{
				AttributeName: aws.String("Created"),
				AttributeType: "N",
			},
			{
				AttributeName: aws.String("Status"),
				AttributeType: "N",
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("lastmodified"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("Id"),
						KeyType:       "HASH",
					},
					{
						AttributeName: aws.String("LastModified"),
						KeyType:       "RANGE",
					},
				},
				Projection: &types.Projection{
					ProjectionType: "ALL",
				},
			},
			{
				IndexName: aws.String("created"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("Id"),
						KeyType:       "HASH",
					},
					{
						AttributeName: aws.String("Created"),
						KeyType:       "RANGE",
					},
				},
				Projection: &types.Projection{
					ProjectionType: "ALL",
				},
			},
			{
				IndexName: aws.String("status"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("Status"),
						KeyType:       "HASH",
					},
					{
						AttributeName: aws.String("LastModified"),
						KeyType:       "RANGE",
					},
				},
				Projection: &types.Projection{
					ProjectionType: "ALL",
				},
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	},
}
