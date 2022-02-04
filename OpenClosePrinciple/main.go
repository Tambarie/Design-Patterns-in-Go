package main

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
	
}

func (f *Filter) FilterByColor(product []Product, color Color) *Product {
	
}

func main()  {
	
}
