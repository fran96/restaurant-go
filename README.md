# restaurant-go

## Overview

restaurant-go is a representation of a restraunt where there a multiple waiters and a kitchen.
It uses GRPC as the architecture to build the API, and it handles creating the order flow.
It uses Kafka as a message queue to retrieve the orders.

## Generate gRPC client and server interfaces

We use protocol buffer compiler `protoc` with a gRPC GO plugin to generate the proto files.
To do so, run the following command:

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    contracts/waiter.proto
```
