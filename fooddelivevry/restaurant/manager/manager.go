package manager

import (
	"strings"
	"sync"

	"github.com/improver2108/golld/fooddelivevry/restaurant"
)

var (
	instance *restaurantManager
	once     sync.Once
)

type restaurantManager struct {
	restaurants []*restaurant.Restaurant
}

func GetInstance() *restaurantManager {
	once.Do(func() {
		instance = &restaurantManager{}
	})
	return instance
}

func (r *restaurantManager) AddRestaurant(restaurant *restaurant.Restaurant) {
	r.restaurants = append(r.restaurants, restaurant)
}

func (r *restaurantManager) SearchByLocation(location string) []*restaurant.Restaurant {
	var res []*restaurant.Restaurant
	for _, rest := range r.restaurants {
		if !strings.EqualFold(rest.GetLocation(), location) {
			continue
		}
		res = append(res, rest)
	}
	return res
}
