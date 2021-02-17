# lb

Range based load balancer. Written while on an airplane with no internet. Think of it as drawing from memory. Meaning it's crap.

# Usage

```bash
# starts the lb
go run cmd/main.go

# Then start two dummy servers
go run cmd/main.go -type dummy -port 8081
go run cmd/main.go -type dummy -port 8082
```


https://blog.yugabyte.com/four-data-sharding-strategies-we-analyzed-in-building-a-distributed-sql-database/

