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

	applyDiscount(&grapes.Product, 20)
	fmt.Println(grapes.Cost)

	products := make([]Product, 5)
	products[0] = Product{101, "Pen", 5, 50, "Stationary"}
	products[1] = Product{105, "Pencil", 2, 100, "Stationary"}
	products[2] = Product{102, "Mouse", 100, 20, "Electronics"}
	products[3] = Product{104, "Marker", 30, 80, "Stationary"}
	products[4] = Product{103, "Charger", 200, 10, "Electronics"}

	fmt.Println("Initial List")
	printProducts(products)

	fmt.Println("After adding stylus")
	stylus := Product{107, "Stylus", 30, 50, "Stationary"}
	addProduct(&products, stylus)
	printProducts(products)

	fmt.Println("Index of stylus = ", indexOf(&products, stylus))
	nonExistingProduct := Product{}
	fmt.Println("Index of a non existing product = ", indexOf(&products, nonExistingProduct))

	fmt.Println("Does the product list include a stylus ? :", includes(&products, stylus))
	costlyProductPredicate := func(product Product) bool {
		return product.Cost >= 100
	}
	costlyProductExists := any(&products, costlyProductPredicate)
	fmt.Println("Are there any costly [cost >= 100] products :", costlyProductExists)
	areAllCostlyProducts := all(&products, costlyProductPredicate)
	fmt.Println("Are all the products costly products ? :", areAllCostlyProducts)

	costlyProducts := filter(&products, costlyProductPredicate)
	fmt.Println("Costly products")
	printProducts(costlyProducts)

	/* costlyStationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary" && product.Cost >= 100
	} */
}

func addProduct(products *[]Product, product Product) {
	*products = append(*products, product)
}

func indexOf(products *[]Product, product Product) int {
	for idx, prod := range *products {
		if prod == product {
			return idx
		}
	}
	return -1
}

func includes(products *[]Product, product Product) bool {
	return indexOf(products, product) != -1
}

func any(products *[]Product, predicate func(Product) bool) bool {
	for _, prod := range *products {
		if predicate(prod) {
			return true
		}
	}
	return false
}

func all(products *[]Product, predicate func(Product) bool) bool {
	for _, prod := range *products {
		if !predicate(prod) {
			return false
		}
	}
	return true
}

func filter(products *[]Product, predicate func(Product) bool) []Product {
	result := []Product{}
	for _, prod := range *products {
		if predicate(prod) {
			result = append(result, prod)
		}
	}
	return result
}

func printProducts(products []Product) {
	for _, product := range products {
		fmt.Println(format(product))
	}
}

func format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func applyDiscount(productPtr *Product, percent float32) {
	productPtr.Cost = productPtr.Cost * ((100 - percent) / 100)
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{Product{id, name, cost, units, category}, expiry}
}
