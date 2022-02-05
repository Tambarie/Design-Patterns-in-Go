package main

import "fmt"

//OCP
// Open for extension, closed for modification
//Specification

type Color int
const  (
	red Color = iota
	green
	blue
)

type Size int
const (
	small Size = iota
	medium 
	large
)

type Product struct {
	name string
	color Color
	size Size
}

type Filter struct {
	//
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product,0)
	for i, product := range products {
		if product.color == color{
			result = append(result,&products[i])
		}
	}

	return result
	
}

func main()  {
	apple := Product{
		name:  "Apple",
		color: green,
		size:  small,
	}

	tree := Product{
		name:  "Tree",
		color: red,
		size:  medium,
	}
	house := Product{
		name:  "House",
		color: blue,
		size:  large,
	}

	products := []Product{apple,tree,house}


	fmt.Printf("Green products (old): \n")
	f := Filter{}

	for _, v := range f.FilterByColor(products,green) {
		fmt.Printf("- %s is green \n",v.name)
	}

}
