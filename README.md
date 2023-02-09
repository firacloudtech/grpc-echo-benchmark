## gRPC-gateway-with-swagger-and-buf

A repository for setting up gRPC, gRPC-Gateway, Buf and benchmarking.

Prerequisites
- [Go](https://golang.org/dl/)
- [Protocol Buffers v3](https://github.com/google/protobuf/releases)
- [gRPC](https://grpc.io/docs/quickstart/go.html)
- [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [Buf](https://buf.build/docs/install)
- make

## Setting up gRPC
Clone the repository.bash
```bash
git clone git@github.com:firacloudtech/grpc-gateway-with-swagger-with-buf.git
```

Run
``` bash
make buf-generate
```

To run the grpc server and grpc-gateway,
Run
``` bash
make run-server
```

To run the grpc client,
Run
``` bash
make run-client
```

To view the swagger documentation, go to http://127.0.0.1:3001/docs/

## sqlc

Update the SQL query in the sql/query.sql and run make sql-generate

For example,

`-- name: CreateProduct :one
INSERT INTO products
(id, name, description, price, category, image_url, created_at, updated_at)
 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
 RETURNING *;`