package main

import (
	"strconv"

	"github.com/firacloudtech/grpc-echo-benchmark/echo-server/handler"
	"github.com/firacloudtech/grpc-echo-benchmark/redis"
	"github.com/labstack/echo/v4"
)

const (
	echoPort = 3002
)

func main() {

	redis.InitRedis()

	e := echo.New()

	h := &handler.Handler{}
	e.POST("/products", h.CreateProduct)

	address := ":" + strconv.Itoa(echoPort)
	e.Logger.Fatal(e.Start(address))
}
