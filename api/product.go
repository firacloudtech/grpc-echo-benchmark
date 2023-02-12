package api

import (
	"context"
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
)

func (p *Product) CreateProduct(c echo.Context, product Product) (string, error) {

	ctx := context.Background()

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
		return "Error", c.String(http.StatusInternalServerError, err.Error())
	}

	if len(rows) > 0 {
		return "Error", c.String(http.StatusConflict, "product already exists, id: "+rows[0].ID)
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
	log.Errorf("index is: %v", result)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
	}

	if err != nil {
		return "Error", c.String(http.StatusInternalServerError, "Unable to save to db: "+err.Error())
	}

	return result.ID, c.JSON(http.StatusCreated, product)
}
