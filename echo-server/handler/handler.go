package handler

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	api "github.com/firacloudtech/grpc-echo-benchmark/api"
)

type (
	Handler struct {
		api.Product
	}
	Product struct {
		api.Product
	}
)

func (i Product) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (m *Product) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

// util

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags root
// @Accept */*
// @Param product body Product true "Product"
// @Produce json
// @Success 200
// @Router /products [post]
func (h *Handler) CreateProduct(c echo.Context) error {
	var p api.Product
	if err := c.Bind(&p); err != nil {
		return err
	}
	id, err := p.CreateProduct(c, p)

	if err != nil {
		fmt.Print("err: " + err.Error())
	}
	fmt.Printf("id: %v", id)

	return err

}
