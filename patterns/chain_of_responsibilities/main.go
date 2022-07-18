package main

import "fmt"

// ChainHandler describes any single handler
// element from the chain of responsibility
type ChainHandler interface {
	Execute(c *Client)
	SetNext(h ChainHandler) ChainHandler
}

// BaseHandler implements setting next element
type BaseHandler struct {
	next ChainHandler
}

func (bh *BaseHandler) SetNext(h ChainHandler) ChainHandler {
	bh.next = h
	return h
}

// AutoResponder is a first element of the chain
type AutoResponder struct {
	BaseHandler
}

func (ar AutoResponder) Execute(c *Client) {
	if c.autoResponderPassed {
		fmt.Println("auto responder passed")
		if ar.next != nil {
			ar.next.Execute(c)
		}
		return
	}
	fmt.Println("if you have problem ... press 1 ...")
	c.autoResponderPassed = true // check this stage as passed
	fmt.Println("auto responder passed")
	if ar.next != nil {
		ar.next.Execute(c)
	}
}

// TechSupportClerk is the second element of the chain
type TechSupportClerk struct {
	BaseHandler
}

func (tsc TechSupportClerk) Execute(c *Client) {
	if c.techSupportOfficeClerkPassed {
		fmt.Println("tech support clerk stage passed")
		if tsc.next != nil {
			tsc.next.Execute(c)
		}
		return
	}
	fmt.Println("Hello, how may I help you ...")
	c.techSupportOfficeClerkPassed = true // check this stage as passed
	fmt.Println("tech support clerk stage passed")
	if tsc.next != nil {
		tsc.next.Execute(c)
	}
}

// TechSupportSpecialist is the second element of the chain
type TechSupportSpecialist struct {
	BaseHandler
}

func (tss TechSupportSpecialist) Execute(c *Client) {
	if c.techSupportSpecialistPassed {
		fmt.Println("tech support specialist stage passed... guess the problem is unsolvable")
		if tss.next != nil {
			tss.next.Execute(c)
		}
		return
	}
	fmt.Println("Let's actually solve your problem")
	c.techSupportSpecialistPassed = true // check this stage as passed
	fmt.Println("tech support specialist stage passed")
	if tss.next != nil {
		tss.next.Execute(c)
	}
}

func main() {
	// init chain of responsibility and set handlers
	chain := new(AutoResponder)
	chain.SetNext(new(TechSupportClerk)).SetNext(new(TechSupportSpecialist))
	// init client
	c := &Client{name: "Peter"}
	// launch processes in the chain
	chain.Execute(c)
}

// Client type describes client of the
// system with all the needed bool checks
type Client struct {
	name                         string
	autoResponderPassed          bool
	techSupportOfficeClerkPassed bool
	techSupportSpecialistPassed  bool
}
