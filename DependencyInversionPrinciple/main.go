package main

import "fmt"

//Dependency Inversion Principle
// KLM should not depend on LLM
//Both should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from *Person
	relationship Relationship
	to *Person
}

// Relationships low-level module
type Relationships  struct {
	relations []Info
}

type RelationshipBrowser interface {
	FindAllChildren(name string) []*Person
}
func (r *Relationships) AddParentAndChild(parent, child *Person)  {
	r.relations = append(r.relations,Info{parent,Parent,child})
	r.relations =append(r.relations,Info{child,Child,parent})
}

// Research high-level module
type Research struct {
	// Breaking the Dependency Injection Principle
	//relationships Relationships
	broswer RelationshipBrowser
}

//func (r *Research) Investigate()  {
//	relations := r.relationships.relations
//	for _, rel := range relations {
//		if rel.from.name == "John" && rel.relationship == Parent{
//			fmt.Println("John has a child called", rel.to.name)
//		}
//	}
//}

func (r *Research) Investigate()  {
	for _, rel := range r.broswer.FindAllChildren("John"){
		fmt.Println("john has a child called" ,rel.name)
	}
}



func (r *Relationships) FindAllChildren (name string)  []*Person{
	result := make([]*Person,0)

	for i, v := range r.relations{
		if v.relationship == Parent && v.from.name == name {
			result = append(result,r.relations[i].to)
		}
	}
return result
}
func main()  {
	parent := Person{name: "John"}
	child1 := Person{name: "Emma"}
	//child2 := Person{name: "Tammy"}

	relationships := Relationships{}

	relationships.AddParentAndChild(&parent,&child1)
	//relationships.AddParentAndChild(&parent,&child2)
	r := &Research{&relationships}
	r.Investigate()
}