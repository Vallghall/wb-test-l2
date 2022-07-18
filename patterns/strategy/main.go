package main

import "fmt"

// Strategy describes wrapper types for
// different algorithms
type Strategy interface {
	Execute(a, b int) int
}

// Addition wraps summing of two numbers
type Addition struct {
}

func (ad Addition) Execute(a, b int) int {
	return a + b
}

// Subtraction wraps subtracting one number from another
type Subtraction struct {
}

func (s Subtraction) Execute(a, b int) int {
	return a - b
}

func main() {
	// init numbers
	a, b := 5, 3
	// init context with Addition strategy
	c := NewContext(new(Addition))
	fmt.Printf("First strategy execution returns %d\n", c.ExecuteStrategy(a, b))
	// change strategy on Subtraction
	c.SetStrategy(new(Subtraction))
	fmt.Printf("Second strategy execution returns %d\n", c.ExecuteStrategy(a, b))
}

// Context works with strategies through common interface
// without knowing about concrete strategies
type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

// SetStrategy changes current strategy with another one
func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

// ExecuteStrategy calls Execute method on given Strategy
func (c Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}
