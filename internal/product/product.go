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
	}
}

// Const for error messages
const (
	ErrorSkuExists     string = "ERROR_SKU_EXISTS"
	ErrorWrongCategory string = "ERROR_WRONG_CATEGORY"
)

// Product Domain
type User struct {
	Id        int
	Name      string
	Sku       string
	Category  ProductCategory
	Price     string
	CreatedAt time.Time
}

// Create a new product instance
func New(n string, s string, c string, p string) (*Product, error) {
	return &Product{
		Name:     n,
		Sku:      s,
		Category: c,
		Price:    p,
	}, nil
}
