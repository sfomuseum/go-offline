GOMOD=vendor

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/create-dynamodb-tables cmd/create-dynamodb-tables/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/add-job cmd/add-job/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/get-job cmd/get-job/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/remove-job cmd/remove-job/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/job-server cmd/job-server/main.go

debug-add-job:
	go run -mod $(GOMOD) -w" cmd/add-job/main.go -instructions '{"hello":"world"}'

# For DynamoDB-related stuff see sfomuseum/go-offline-gocloud
