package main

import (
	"context"
	"flag"
	aa_dynamodb "github.com/aaronland/go-aws-dynamodb"
	sfom_dynamodb "github.com/sfomuseum/go-offline/dynamodb"
	"log"
)

func main() {

	client_uri := flag.String("client-uri", "", "...")
	refresh := flag.Bool("refresh", false, "...")

	flag.Parse()

	ctx := context.Background()

	client, err := sfom_dynamodb.NewClientWithURI(ctx, *client_uri)

	if err != nil {
		log.Fatalf("Failed to create client, %v", err)
	}

	table_opts := &aa_dynamodb.CreateTablesOptions{
		Tables:  sfom_dynamodb.DynamoDBTables,
		Refresh: *refresh,
	}

	err = aa_dynamodb.CreateTables(client, table_opts)

	if err != nil {
		log.Fatalf("Failed to create tables, %v", err)
	}
}
