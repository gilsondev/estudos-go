package entity

import (
	"errors"
	"time"

	"github.com/gilsondev/estudos-go/go-expert-apis/pkg/entity"
)

var (
	ErrorIDIsRequired    = errors.New("id is required")
	ErrorNameIsRequired  = errors.New("name is required")
	ErrorPriceIsRequired = errors.New("price is required")
	ErrorInvalidID       = errors.New("invalid id")
	ErrorInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorInvalidID
	}
	if p.Name == "" {
		return ErrorNameIsRequired
	}
	if p.Price == 0 {
		return ErrorPriceIsRequired
	}
	if p.Price <= 0 {
		return ErrorInvalidPrice
	}
	return nil
}
