package mocks

import (
	"time"

	"github.com/sebajax/go-vertical-slice-architecture/internal/product"
)

type mockProductRepository struct{}

func NewMockProductRepository() product.ProductRepository {
	return &mockProductRepository{}
}

func (mock *mockProductRepository) Save(p *product.Product) (int64, error) {
	return 1, nil
}

func (mock *mockProductRepository) GetBySku(email string) (*product.Product, bool, error) {
	return &product.Product{
		Id:        1,
		Name:      "ps5",
		Sku:       "123456sony",
		Category:  product.SmartWatch,
		Price:     300.00,
		CreatedAt: time.Now(),
	}, true, nil
}
