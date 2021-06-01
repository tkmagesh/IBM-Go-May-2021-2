package main

import (
	"fmt"
	"methods-demo/entities"
)

func main() {
	//var product Product
	//product := Product{}
	//p1 := Product{100, "Pen", 10, 100, "Stationary"}
	p1 := entities.Product{Id: 100, Name: "Pen", Cost: 10, Units: 100, Category: "Stationary"}
	fmt.Println(p1)
	p1.ApplyDiscount(10)
	fmt.Printf("After applying 10%% discount, product = %v\n", p1)

	/* grapes := PerishableProduct{product: Product{102, "Grapes", 60, 50, "Food"}, Expiry: "2 Days"} */
	/* grapes := PerishableProduct{Product{102, "Grapes", 60, 50, "Food"}, "2 Days"} */
	/* fmt.Println(grapes)
	fmt.Println(grapes.product.Cost) */

	//grapes := PerishableProduct{Product{102, "Grapes", 60, 50, "Food"}, "2 Days"}
	//fmt.Println(grapes.Product.Cost)
	//using the NewPerishableProduct utility function
	grapes := entities.NewPerishableProduct(102, "Grapes", 60, 50, "Food", "2 Days")
	fmt.Println(grapes.Cost)

	grapes.ApplyDiscount(20)
	fmt.Println(grapes.Cost)

	productsSlice := make([]entities.Product, 5)
	productsSlice[0] = entities.Product{101, "Pen", 5, 50, "Stationary"}
	productsSlice[1] = entities.Product{103, "Pencil", 2, 100, "Stationary"}
	productsSlice[2] = entities.Product{102, "Mouse", 100, 20, "Electronics"}
	productsSlice[3] = entities.Product{104, "Marker", 30, 80, "Stationary"}
	productsSlice[4] = entities.Product{105, "Charger", 200, 10, "Electronics"}

	products := entities.Products(productsSlice)
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("After adding stylus")
	stylus := entities.Product{107, "Stylus", 30, 50, "Stationary"}
	products.AddProduct(stylus)
	fmt.Println(products)

	fmt.Println("Index of stylus = ", products.IndexOf(stylus))
	nonExistingProduct := entities.Product{}
	fmt.Println("Index of a non existing product = ", products.IndexOf(nonExistingProduct))

	fmt.Println("Does the product list include a stylus ? :", products.Includes(stylus))
	costlyProductPredicate := func(product entities.Product) bool {
		return product.Cost >= 100
	}
	costlyProductExists := products.Any(costlyProductPredicate)
	fmt.Println("Are there any costly [cost >= 100] products :", costlyProductExists)
	areAllCostlyProducts := products.All(costlyProductPredicate)
	fmt.Println("Are all the products costly products ? :", areAllCostlyProducts)

	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println("Costly products")
	fmt.Println(costlyProducts)

	/* costlyStationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary" && product.Cost >= 100
	} */

	fmt.Println()
	fmt.Println("Default List")
	fmt.Println(products)

	fmt.Println("After sorting (default)")
	products.Sort()
	fmt.Println(products)

	fmt.Println("After sorting by name")
	products.SortByName()
	fmt.Println(products)

	fmt.Println("After sorting by cost")
	products.SortByCost()
	fmt.Println(products)

	fmt.Println("After sorting by units")
	products.SortByUnits()
	fmt.Println(products)

	fmt.Println("After sorting by category")
	products.SortByCategory()
	fmt.Println(products)
}
