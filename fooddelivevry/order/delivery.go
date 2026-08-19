package order

type DeliveryOrder struct {
	*baseOrder
	address string
}

func NewDeliveryOrder(order *baseOrder, address string) *DeliveryOrder {
	return &DeliveryOrder{
		baseOrder: order,
		address:   address,
	}
}

func (o *DeliveryOrder) GetType() string {
	return "delivery"
}

func (o *DeliveryOrder) GetAddress() string {
	return o.address
}
