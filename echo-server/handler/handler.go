package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/firacloudtech/grpc-echo-benchmark/redis"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

// Handlers
func (h *Handler) CreateProduct(c echo.Context) error {
	product := new(Product)

	ctx := context.Background()

	if err := c.Bind(&product); err != nil {
		return c.String(http.StatusBadRequest, ("bad request: " + err.Error()))
	}

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result := redis.Client.SScan(ctx, "product", 0, product).Result()

	log.Infof("index is: %v", result)

	if result >= 0 {
		return c.String(http.StatusBadRequest, "Item already exists: "+err2.Error())
	}

	_, err := redis.Client.RPush(ctx, "products",
		product).Result()

	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to save to redis: "+err.Error())
	}

	return c.JSON(http.StatusCreated, product)

}
