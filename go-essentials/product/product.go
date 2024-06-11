package product

type Product struct {
	ID    string
	Title string
	Price float64
}

func New(id string, title string, price float64) *Product {
	return &Product{
		ID:    id,
		Title: title,
		Price: price,
	}
}
