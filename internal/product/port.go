package product

// Product port interface definition for depedency injection
type ProductRepository interface {
	Save(p *Product) (int64, error)
	GetBySku(sku string) (*Product, bool, error)
}
