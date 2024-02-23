package service

import (
	"log"

	"github.com/sebajax/go-vertical-slice-architecture/internal/product"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/apperror"
)

// Product use cases (port injection)
type CreateProductService struct {
	productRepository product.ProductRepository
}

// Create a new product service use case instance
func NewCreateProductService(repository product.ProductRepository) *CreateProductService {
	// return the pointer to product service
	return &CreateProductService{
		productRepository: repository,
	}
}

// Create a new product and store the product into the database
func (service *CreateProductService) CreateProduct(p *product.Product) (int64, error) {
	_, check, err := service.productRepository.GetBySku(p.Sku)
	// check if product sky does not exist and no database error ocurred
	if err != nil {
		// database error
		log.Println(err)
		err := apperror.InternalServerError()
		return 0, err
	}
	if check {
		// product sku found
		log.Println(p, product.ErrorSkuExists)
		err := apperror.BadRequest(product.ErrorSkuExists)
		return 0, err
	}

	// create the new product and return the id
	id, err := service.productRepository.Save(p)
	if err != nil {
		// database error
		log.Println(err)
		err := apperror.InternalServerError()
		return 0, err
	}

	// product created successfuly
	return id, nil
}
