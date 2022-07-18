package main

import "fmt"

// Animal is an interface in which we have to add functionality
type Animal interface {
	Voice()
}

// Visitor describes methods for visiting types of interest
type Visitor interface {
	visitForDog(d *Dog) // method for visiting dog
	visitForCat(c *Cat) // method for visiting cat
}

// TrickBag is a visitor adding tricks functionality to the Dog type
type TrickBag struct {
	// ...
}

// visitForDog adds functionality to Dog type
func (b TrickBag) visitForDog(d *Dog) {
	fmt.Printf("Vigorous %s does new tricks\n", d.name)
}

// visitForCat adds functionality to Cat type
func (b TrickBag) visitForCat(c *Cat) {
	fmt.Printf("Cute %s does new tricks\n", c.name)
}

// Dog type represents a dog
type Dog struct {
	name string
}

func (d Dog) Voice() {
	fmt.Println("woof")
}

func NewDog(name string) *Dog {
	return &Dog{name}
}

// Accept is a method for accepting functionality from visitors
func (d *Dog) Accept(v Visitor) {
	v.visitForDog(d)
}

// Cat type represents a cat
type Cat struct {
	name string
}

func (c Cat) Voice() {
	fmt.Println("meow")
}

func NewCat(name string) *Cat {
	return &Cat{name: name}
}

// Accept is a method for accepting functionality from visitors
func (c *Cat) Accept(v Visitor) {
	v.visitForCat(c)
}

func main() {
	jeffy := NewDog("Jeffy") // init Dog by the name Jeffy
	puss := NewCat("Puss")   // init Dog by the name Puss

	visitor := new(TrickBag)
	jeffy.Accept(visitor) // learn Jeffy soma new tricks via TrickBag visitor
	puss.Accept(visitor)
}
