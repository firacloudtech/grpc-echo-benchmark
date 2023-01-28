module github.com/firacloudtech/grpc-echo-benchmark

go 1.19

require google.golang.org/protobuf v1.28.1 // indirect

require (
	buf.build/gen/go/grpc-ecosystem/grpc-gateway/protocolbuffers/go v1.28.1-20221127060915-a1ecdc58eccd.4 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.3.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
)

require (
	buf.build/gen/go/firacloudtech/grpc-echo-benchmark/grpc/go v1.2.0-20230122134124-c7efdc22f76a.4
	buf.build/gen/go/firacloudtech/grpc-echo-benchmark/protocolbuffers/go v1.28.1-20230122134124-c7efdc22f76a.4
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	google.golang.org/grpc v1.51.0
)
