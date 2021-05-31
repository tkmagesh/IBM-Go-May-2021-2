package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/* type PerishableProduct struct {
	product Product
	Expiry  string
} */

type PerishableProduct struct {
	Product
	Expiry string
}

func main() {
	//var product Product
	//product := Product{}
	//p1 := Product{100, "Pen", 10, 100, "Stationary"}
	p1 := Product{Id: 100, Name: "Pen", Cost: 10, Units: 100, Category: "Stationary"}
	fmt.Println(p1)
	applyDiscount(&p1, 10)
	fmt.Printf("After applying 10%% discount, product = %v\n", p1)

	/* grapes := PerishableProduct{product: Product{102, "Grapes", 60, 50, "Food"}, Expiry: "2 Days"} */
	/* grapes := PerishableProduct{Product{102, "Grapes", 60, 50, "Food"}, "2 Days"} */
	/* fmt.Println(grapes)
	fmt.Println(grapes.product.Cost) */

	//grapes := PerishableProduct{Product{102, "Grapes", 60, 50, "Food"}, "2 Days"}
	//fmt.Println(grapes.Product.Cost)
	//using the NewPerishableProduct utility function
	grapes := NewPerishableProduct(102, "Grapes", 60, 50, "Food", "2 Days")
	fmt.Println(grapes.Cost)
}

func applyDiscount(productPtr *Product, percent float32) {
	productPtr.Cost = productPtr.Cost * ((100 - percent) / 100)
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{Product{id, name, cost, units, category}, expiry}
}
