package main

import (
	"strconv"

	"github.com/firacloudtech/grpc-echo-benchmark/db"
	"github.com/firacloudtech/grpc-echo-benchmark/echo-server/handler"
	"github.com/labstack/echo/v4"
)

const (
	echoPort = 3002
)

func main() {

	db.InitDB()
	defer db.Db.Close()

	e := echo.New()

	h := &handler.Handler{}
	e.POST("/products", h.CreateProduct)

	address := ":" + strconv.Itoa(echoPort)
	e.Logger.Fatal(e.Start(address))
}
