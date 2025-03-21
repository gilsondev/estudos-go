package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestProductWhenNameIsInvalid(t *testing.T) {
	p, err := NewProduct("", 10.0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrorNameIsRequired, err)
	assert.Nil(t, p)
}

func TestProductWhenPriceIsZero(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrorPriceIsRequired, err)
	assert.Nil(t, p)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -10.0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrorInvalidPrice, err)
	assert.Nil(t, p)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
