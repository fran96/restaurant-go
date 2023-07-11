# restaurant-go

## Overview

restaurant-go is a representation of a restraunt where there are 2 independent hosted services, `waiter` and `kitchen`.
It uses GRPC as the transport for sync API requests between the two services
Kafka is used as an event stream between them for asynchrnous communication to retrieve the orders status once completed.

Orders
1. An order is sent throught GRPC request from the client entrypoint.
2. The order request is forwarded by the waiter to the kitchen.
3. The kitchen replies saying whether the order was accepted with a generated orderID. The waiter returns this to the client.
4. Once kitchen is ready (random sleep), it produces a Kafka message as a completed event.
5. The waiter consumes the Kafka message and prints out to console.

## Generate gRPC client and server interfaces

We use protocol buffer compiler `protoc` with a gRPC GO plugin to generate the proto files.
To do so, run the following command:

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    contracts/waiter.proto
```

## Run waiter and kitchen services
In one terminal run `go run -tags dynamic internal/waiter/cmd/waiter/main.go` and in another terminal run `go run -tags dynamic internal/kitchen/cmd/kitchen/main.go`.
You must also run `docker-compose up -d` to run zookeeper, schema-registry and Kafka.
