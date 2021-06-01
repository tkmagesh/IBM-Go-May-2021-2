package entities

import (
	"fmt"
	"sort"
)

type Products []Product

func (products *Products) AddProduct(product Product) {
	*products = append(*products, product)
}

func (products *Products) IndexOf(product Product) int {
	for idx, prod := range *products {
		if prod == product {
			return idx
		}
	}
	return -1
}

func (products *Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products *Products) Any(predicate func(Product) bool) bool {
	for _, prod := range *products {
		if predicate(prod) {
			return true
		}
	}
	return false
}

func (products *Products) All(predicate func(Product) bool) bool {
	for _, prod := range *products {
		if !predicate(prod) {
			return false
		}
	}
	return true
}

func (products *Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, prod := range *products {
		if predicate(prod) {
			result = append(result, prod)
		}
	}
	return result
}

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%v\n", product)
	}
	return result
}

/* methods of 'Interface' interface */

func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Sort() {
	sort.Sort(products)
}

type ByName struct {
	Products
}

func (list ByName) Less(i, j int) bool {
	return list.Products[i].Name < list.Products[j].Name
}

type ByCost struct {
	Products
}

func (list ByCost) Less(i, j int) bool {
	return list.Products[i].Cost < list.Products[j].Cost
}

type ByUnits struct {
	Products
}

func (list ByUnits) Less(i, j int) bool {
	return list.Products[i].Units < list.Products[j].Units
}

type ByCategory struct {
	Products
}

func (list ByCategory) Less(i, j int) bool {
	return list.Products[i].Category < list.Products[j].Category
}

func (products Products) SortBy(attrName string) {
	switch attrName {
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Category":
		sort.Sort(ByCategory{products})
	}
}
