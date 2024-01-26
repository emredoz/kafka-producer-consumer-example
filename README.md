# kafka-producer-consumer-example

### Requirements
- go1.21
- kafka


### Generate Proto files
```
$ protoc --proto_path=. --go_out=./common  ./common/model/proto/notification.proto
```

### Send Notification
```
$ curl --location 'http://localhost:8081/api/v1/notification' \
--header 'Content-Type: application/json' \
--data '{
    "to":  "dahi",
    "from":"emre",
    "message":"hello"
}'
```

### Run producer
```
$ go mod tidy
$ go build -o .producer-app producer/main.go
$ ./.producer-app  
```

### Run consumer
```
$ go mod tidy
$ go build -o .consumer-app consumer/main.go
$ ./.consumer-app  
```


