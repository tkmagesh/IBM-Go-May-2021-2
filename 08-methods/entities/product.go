package entities

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (productPtr *Product) ApplyDiscount(percent float32) {
	productPtr.Cost = productPtr.Cost * ((100 - percent) / 100)
}
