package service

import (
	"errors"
	"testing"

	"github.com/sebajax/go-vertical-slice-architecture/internal/product"
	"github.com/sebajax/go-vertical-slice-architecture/internal/product/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductService_CreateProduct_Success(t *testing.T) {
	var insertedId int64 = 1
	// Create a new Product
	p, _ := product.New("test_product", "test_sku", product.Laptop, 1200)

	// Create a mock ProductRepository instance
	mockProductRepository := &mock.ProductRepository{}

	// Set expectations on the mock
	mockProductRepository.EXPECT().GetBySku(p.Sku).Return(nil, false, nil)
	mockProductRepository.EXPECT().Save(p).Return(insertedId, nil)

	// Pass the mock as a dependency
	productService := NewCreateProductService(mockProductRepository)

	// Call the method being tested
	id, err := productService.CreateProduct(p)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, insertedId, id)

	// Assert that the method was called exactly once or not called
	mockProductRepository.AssertCalled(t, "GetBySku", p.Sku)
	mockProductRepository.AssertCalled(t, "Save", p)

	// Assert that the expectations were met
	mockProductRepository.AssertExpectations(t)
}

func TestCreateProductService_CreateProduct_GetBySku_Failure(t *testing.T) {
	// Create a new Product
	p, _ := product.New("test_product", "test_sku", product.Laptop, 1200)

	// Create a mock ProductRepository instance
	mockProductRepository := &mock.ProductRepository{}

	// Set expectations on the mock
	mockProductRepository.EXPECT().GetBySku(p.Sku).Return(p, true, nil)

	// Pass the mock as a dependency
	productService := NewCreateProductService(mockProductRepository)

	// Call the method being tested
	id, err := productService.CreateProduct(p)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, int64(0), id)

	// Assert that the method was called exactly once or not called
	mockProductRepository.AssertCalled(t, "GetBySku", p.Sku)
	mockProductRepository.AssertNotCalled(t, "Save")

	// Assert that the expectations were met
	mockProductRepository.AssertExpectations(t)
}

func TestCreateProductService_CreateProduct_Save_Failure(t *testing.T) {
	// Create a new Product
	p, _ := product.New("test_product", "test_sku", product.Laptop, 1200)

	// Create a mock ProductRepository instance
	mockProductRepository := &mock.ProductRepository{}

	// Set expectations on the mock
	mockProductRepository.EXPECT().GetBySku(p.Sku).Return(nil, false, nil)
	mockProductRepository.EXPECT().Save(p).Return(0, errors.New("DB ERROR"))

	// Pass the mock as a dependency
	productService := NewCreateProductService(mockProductRepository)

	// Call the method being tested
	id, err := productService.CreateProduct(p)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, int64(0), id)

	// Assert that the method was called exactly once or not called
	mockProductRepository.AssertCalled(t, "GetBySku", p.Sku)
	mockProductRepository.AssertCalled(t, "Save", p)

	// Assert that the expectations were met
	mockProductRepository.AssertExpectations(t)
}
