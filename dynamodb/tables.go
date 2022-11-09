package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDBTables is a map whose keys are DynamoDB table names and whose values are `dynamodb.CreateTableInput` instances.
var DynamoDBTables = map[string]*dynamodb.CreateTableInput{
	"offlinejobs": &dynamodb.CreateTableInput{
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"), // partition key
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("LastModified"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Created"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Status"),
				AttributeType: aws.String("N"),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("lastmodified"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("Id"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("LastModified"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
			{
				IndexName: aws.String("created"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("Id"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("Created"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
			{
				IndexName: aws.String("status"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("Status"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("LastModified"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
		},
		BillingMode: aws.String("PAY_PER_REQUEST"),
		// TableName:   set inline below
	},
}
