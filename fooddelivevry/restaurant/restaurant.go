package restaurant

import "github.com/improver2108/golld/fooddelivevry/menuitem"

type Restaurant struct {
	id       int
	name     string
	location string
	menu     []*menuitem.MenuItem
}

func Init(name, location string) *Restaurant {
	return &Restaurant{name: name, location: location}
}

func (r *Restaurant) GetName() string {
	return r.name
}

func (r *Restaurant) SetName(name string) {
	r.name = name
}

func (r *Restaurant) GetLocation() string {
	return r.location
}

func (r *Restaurant) SetLocation(location string) {
	r.location = location
}

func (r *Restaurant) GetMenu() []*menuitem.MenuItem {
	return r.menu
}

func (r *Restaurant) AddMenuItem(m *menuitem.MenuItem) {
	r.menu = append(r.menu, m)
}
