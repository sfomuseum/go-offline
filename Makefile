GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/add-job cmd/add-job/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/get-job cmd/get-job/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/get-job cmd/schedule-job/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/remove-job cmd/remove-job/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/job-server cmd/job-server/main.go

debug-add-job:
	go run -mod $(GOMOD) -w" cmd/add-job/main.go -instructions '{"hello":"world"}'

# For DynamoDB-related stuff see sfomuseum/go-offline-gocloud

lambda:
	@make lambda-server

lambda-server:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f server.zip; then rm -f server.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags lambda.norpc -o bootstrap cmd/job-server/main.go
	zip server.zip bootstrap
	rm -f bootstrap

debug-server:
	go run cmd/job-server/main.go \
		-offline-database-uri syncmap:// \
		-offline-queue-uri '*=slog://' \
		-authenticator-uri sharedsecret://s33kret

debug-server-queue-config:
	go run cmd/job-server/main.go \
		-offline-queue-config-uri $(CONFIG_URI) \
		-authenticator-uri sharedsecret://s33kret
