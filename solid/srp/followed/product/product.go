package product

type Product struct {
	Name  string
	Price float32
}

func Init(name string, price float32) *Product {
	return &Product{name, price}
}
