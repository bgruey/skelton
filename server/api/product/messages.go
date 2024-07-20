package product

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Products []Product
