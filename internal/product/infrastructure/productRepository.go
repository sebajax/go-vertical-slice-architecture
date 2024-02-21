package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/sebajax/go-vertical-slice-architecture/internal/product"
	"github.com/sebajax/go-vertical-slice-architecture/pkg/database"
)

// Product repository for querying the database
type productRepository struct {
	db *database.DbConn
}

// Create a product instance repository
func NewProductRepository(dbcon *database.DbConn) product.ProductRepository {
	return &productRepository{db: dbcon}
}

// Stores a new product in the database
func (repo *productRepository) Save(p *product.Product) (int64, error) {
	// Get the id inserted in the database
	var id int64

	query := `INSERT INTO product (name, sku, category, price) 
				VALUES ($1, $2, $3, $4) RETURNING id`
	err := repo.db.DbPool.QueryRow(query, p.Name, p.Sku, p.Category, p.Price).Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("id: ", id)

	// No errors return the product id inserted
	return id, nil
}

// Gets the product by the sku
func (repo *productRepository) GetBySku(sku string) (*product.Product, bool, error) {
	p := product.Product{}
	var category string
	query := `SELECT id, name, sku, category, price, created_at
				FROM product 
				WHERE sku = $1`
	err := repo.db.DbPool.QueryRow(query, sku).Scan(&p.Id, &p.Name, &p.Sku, &category, &p.Price, &p.CreatedAt)
	if err != nil {
		// Not found, but not an error
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		// An actual error occurred
		return nil, false, err
	}

	// Parse to Enum category
	p.Category = product.ParseProductCategory(category)

	// Found the item
	return &p, true, nil
}
