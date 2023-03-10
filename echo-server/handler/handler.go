package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/firacloudtech/grpc-echo-benchmark/db"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type (
	Product struct {
		ID          string    `json:"id,omitempty"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
		Category    string    `json:"category"`
		ImageURL    string    `json:"image_url"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	Handler struct {
		Product
	}
)

// util

func (i Product) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (m *Product) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

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
	product := new(Product)

	ctx := context.Background()

	if err := c.Bind(&product); err != nil {
		return c.String(http.StatusBadRequest, ("bad request: " + err.Error()))
	}

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	queries := db.New(db.Db)

	rows, err := queries.ListProducts(ctx, db.ListProductsParams{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if len(rows) > 0 {
		return c.String(http.StatusConflict, "product already exists, id: "+rows[0].ID)
	}

	id := uuid.New()

	params := db.CreateProductParams{
		ID:          id.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		ImageUrl:    product.ImageURL,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	result, err := queries.CreateProduct(ctx, params)
	log.Infof("index is: %v", result)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to save to db: "+err.Error())
	}

	return c.JSON(http.StatusCreated, product)

}
