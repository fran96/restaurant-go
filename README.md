# restaurant-go

## Overview

restaurant-go is a representation of a restraunt where there are 2 independent hosted services, `waiter` and `kitchen`.
It uses GRPC as the transport for sync API requests between the two services
Kafka is used as an event stream between them for asynchronous communication to retrieve the orders status once completed.

<img width="1042" alt="Screenshot 2023-07-12 at 09 29 13" src="https://github.com/fran96/restaurant-go/assets/20288513/073900eb-8af0-480c-a7f7-f306b2b1129e">


### Orders
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
- In one terminal run `go run -tags dynamic internal/waiter/cmd/waiter/main.go` and in another terminal run `go run -tags dynamic internal/kitchen/cmd/kitchen/main.go`.
- You must also run `docker-compose up -d` to run zookeeper, schema-registry and Kafka.
- The order request can then be triggered through BloomRPC.


<img width="1023" alt="Screenshot 2023-07-11 at 19 00 20" src="https://github.com/fran96/restaurant-go/assets/20288513/06c89bbb-6547-4f7f-87e7-e65deec76c41">

*** 

<img width="1030" alt="Screenshot 2023-07-11 at 19 00 09" src="https://github.com/fran96/restaurant-go/assets/20288513/c55eda07-58bf-4b7c-96ad-1724492a8b54">






