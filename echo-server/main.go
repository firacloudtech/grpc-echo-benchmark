package main

import (
	"strconv"

	"github.com/firacloudtech/grpc-echo-benchmark/db"
	"github.com/firacloudtech/grpc-echo-benchmark/echo-server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/firacloudtech/grpc-echo-benchmark/echo-server/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	echoPort = 3002
)

// @title Simple Echo Service
// @version 1.0
// @description This is a simple echo server
// @license.name Apache 2.0
// @host localhost:3002
// @BasePath /
// @schemas http
func main() {

	db.InitDB()
	defer db.Db.Close()

	// initiate echo server
	e := echo.New()
	h := &handler.Handler{}

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.POST("/products", h.CreateProduct)

	address := ":" + strconv.Itoa(echoPort)
	e.Logger.Fatal(e.Start(address))
}
