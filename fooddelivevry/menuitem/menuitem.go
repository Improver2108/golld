package menuitem

type MenuItem struct {
	code  string
	name  string
	price float32
}

func Init(code string, name string, price float32) *MenuItem {
	return &MenuItem{code: code, name: name, price: price}
}

func (i *MenuItem) GetCode() string {
	return i.code
}

func (i *MenuItem) GetName() string {
	return i.name
}

func (i *MenuItem) GetPrice() float32 {
	return i.price
}

func (i *MenuItem) SetName(name string) {
	i.name = name
}

func (i *MenuItem) SetPrice(price float32) {
	i.price = price
}
