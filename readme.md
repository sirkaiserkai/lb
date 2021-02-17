# lb

Range based load balancer.

# Usage

```bash
# starts the lb
go run cmd/main.go

# Then start two dummy servers
go run cmd/main.go -type dummy -port 8081
go run cmd/main.go -type dummy -port 8082
```


https://blog.yugabyte.com/four-data-sharding-strategies-we-analyzed-in-building-a-distributed-sql-database/