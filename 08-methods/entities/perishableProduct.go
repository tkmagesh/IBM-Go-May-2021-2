package entities

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{Product{id, name, cost, units, category}, expiry}
}

func (productPtr *PerishableProduct) ApplyDiscount(percent float32) {
	productPtr.Cost = productPtr.Cost * ((100 - percent) / 100)
}
