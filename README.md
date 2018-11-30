# Mysql client to execute simple SQL queries
MySQL client to execute simple SQL queries (without installing mysql-client package)

# Build

## Requirements

1. docker-ce

```
$ docker build -t mysql_client .
$ docker run --rm -it mysql_client -host localhost -user password -password password  -database mysql -query "show tables"
```

# Retreive artifacts to local folder

```
$ docker run --rm -w /go/src/app -v $(pwd):/app:rw --entrypoint /bin/bash -it mysql_client -c "cp app /app/mysql_client"
$ ./mysql_client -h
Usage of ./mysql_client:
  -database string
    	Database to connect
  -host string
    	The address of MySQL server (default "localhost")
  -password string
    	Password for MySQL user
  -port int
    	Port of MySQL service (default 3306)
  -query string
    	SQL query to execute
  -user string
    	MySQL user (default "root")
```
