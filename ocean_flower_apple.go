package personalcare

import (
	"fmt"
	"strings"
)

// Product defines a product with a name and its ingredients
type Product struct {
	name       string
	ingredients []string
}

// NewProduct returns a new Product object
func NewProduct(name string, ingredients []string) *Product {
	return &Product{
		name:       name,
		ingredients: ingredients,
	}
}

// String returns a nicely formatted representation of the Product object
func (p *Product) String() string {
	return fmt.Sprintf("name: %sIngredients: %s",
		p.name, strings.Join(p.ingredients, ", "))
}

// NaturalProducts is a slice of Product objects
type NaturalProducts []*Product

// Len returns the length of a NaturalProducts slice
func (p NaturalProducts) Len() int {
	return len(p)
}

// Swap swaps two elements within a NaturalProducts slice
func (p NaturalProducts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less returns whether or not the element at index i is less than the element at index j
func (p NaturalProducts) Less(i, j int) bool {
	return p[i].name < p[j].name
}

// ByIngredients is a type for sorting NaturalProducts by the number of ingredients
type ByIngredients struct {
	NaturalProducts
}

// Less returns whether or not the element at index i has fewer ingredients than the element at index j
func (p ByIngredients) Less(i, j int) bool {
	return len(p.NaturalProducts[i].ingredients) < len(p.NaturalProducts[j].ingredients)
}

// NaturalProductsCollection is a collection of NaturalProducts
type NaturalProductsCollection struct {
	products NaturalProducts
}

// Add adds a product to the collection
func (c *NaturalProductsCollection) Add(p *Product) {
	c.products = append(c.products, p)
}

// Sort sorts the collection by name
func (c *NaturalProductsCollection) SortByName() {
	sort.Sort(c.products)
}

// Sort sorts the collection by ingredients
func (c *NaturalProductsCollection) SortByIngredients() {
	sort.Sort(ByIngredients{c.products})
}

// String returns a nicely formatted string of all the products in the collection
func (c *NaturalProductsCollection) String() string {
	products := make([]string, 0, len(c.products))
	for _, p := range c.products {
		products = append(products, p.String())
	}
	return strings.Join(products, "\n")
}

// NaturalProductsApp implements a program to view and manage natural products
type NaturalProductsApp struct {
}

// Run executes the app
func (app *NaturalProductsApp) Run() {
	collection := &NaturalProductsCollection{
		products: NaturalProducts{
			NewProduct("Coconut and Vanilla Body Scrub", []string{"Coconut Oil", "Vanilla Extract"}),
			NewProduct("Lavender and Rosemary Diffuser Oil", []string{"Lavender Oil", "Rosemary Oil"}),
		},
	}
	fmt.Println(collection)
	fmt.Println("Sorting products by ingredient count...")
	collection.SortByIngredients()
	fmt.Println(collection)
}