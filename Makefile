cli:
	go build -mod vendor -o bin/create-dynamodb-tables cmd/create-dynamodb-tables/main.go
	go build -mod vendor -o bin/add-job cmd/add-job/main.go
	go build -mod vendor -o bin/get-job cmd/get-job/main.go
	go build -mod vendor -o bin/remove-job cmd/remove-job/main.go
	go build -mod vendor -o bin/job-status-server cmd/job-status-server/main.go
