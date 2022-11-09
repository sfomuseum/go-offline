# go-offline

Minimalist and opinionated (read: What is the simplest and dumbest) offline task (job) coordination system.

_Work in progress._

## Documentation

Documentation is incomplete at this time.

## Tools

### add-job

```
$> go run -mod vendor cmd/add-job/main.go -instructions '{"hello":"world"}'
1590159022869188608
```

### get-job

```
$> go run -mod vendor cmd/get-job/main.go -job-id 1590159022869188608
{"id":1590159022869188608,"status":0,"created":1667958429,"lastmodified":1667958429,"instruction":{"hello":"world"}}
```

### remove-job

```
$> go run -mod vendor cmd/remove-job/main.go -job-id 1590159022869188608

$> go run -mod vendor cmd/get-job/main.go -job-id 1590159022869188608
2022/11/08 17:47:45 Failed to add job, Failed to get job, Not found
exit status 1
```

