package product

import "time"

// ProductCategory represents the categories of electronic products
type ProductCategory int

// Enumeration of product categories
const (
	Laptop ProductCategory = iota
	Smartphone
	Tablet
	SmartWatch
	Headphones
	Camera
	Television
	Other
)

// String representation of the ProductCategory
func (p ProductCategory) String() string {
	return [...]string{
		"Laptop",
		"Smartphone",
		"Tablet",
		"SmartWatch",
		"Headphones",
		"Camera",
		"Television",
		"Other",
	}[p]
}

// Const for error messages
const (
	ErrorSkuExists     string = "ERROR_SKU_EXISTS"
	ErrorWrongCategory string = "ERROR_WRONG_CATEGORY"
)

// Product Domain
type Product struct {
	Id        int64
	Name      string
	Sku       string
	Category  ProductCategory
	Price     float64
	CreatedAt time.Time
}

// Create a new product instance
func New(n string, s string, c ProductCategory, p float64) (*Product, error) {
	return &Product{
		Name:     n,
		Sku:      s,
		Category: c,
		Price:    p,
	}, nil
}
