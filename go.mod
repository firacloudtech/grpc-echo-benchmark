module github.com/firacloudtech/grpc-echo-benchmark

go 1.19

require google.golang.org/protobuf v1.28.1

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.2.0 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
)

require (
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0
	github.com/labstack/echo/v4 v4.10.0
	github.com/labstack/gommon v0.4.0
	github.com/lib/pq v1.10.7
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37
	google.golang.org/grpc v1.51.0
)

replace github.com/firacloudtech/grpc-echo-benchmark/echo-server/handler v0.0.0 => ../echo-server/handler

replace github.com/firacloudtech/grpc-echo-benchmark/db v0.0.0 => ../db
