package order

type PickupOrder struct {
	*baseOrder
	restaurantAddress string
}

func NewPickupOrder(order *baseOrder, address string) *PickupOrder {
	return &PickupOrder{
		baseOrder:         order,
		restaurantAddress: address,
	}
}

func (o *PickupOrder) GetType() string {
	return "pickup"
}

func (o *PickupOrder) GetAddress() string {
	return o.restaurantAddress
}
