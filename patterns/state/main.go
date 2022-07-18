package main

import (
	"errors"
	"fmt"
	"log"
)

// State describes vending machine states
type State interface {
	RequestItem() error
	InsertMoney(amount float64) error
	DispenseItem() error
	AddItem(n int) error
}

// HasItemState represents vending machine's
// state when it has items
type HasItemState struct {
	vm *VendingMachine
}

func (h *HasItemState) RequestItem() error {
	fmt.Println("Requesting item...")
	h.vm.state = &PaymentState{h.vm, 0}
	return nil
}

func (h *HasItemState) InsertMoney(money float64) error {
	return errors.New("request an item first")
}

func (h *HasItemState) DispenseItem() error {
	return errors.New("no item to take until payment")
}

func (h *HasItemState) AddItem(n int) error {
	return errors.New("still has items left")
}

// PaymentState represents
type PaymentState struct {
	vm             *VendingMachine
	insertedAmount int
}

func (p *PaymentState) RequestItem() error {
	return errors.New("payment required")
}

func (p *PaymentState) InsertMoney(amount float64) error {
	p.insertedAmount += int(amount * 100)
	if p.insertedAmount == int(amount*100) {
		fmt.Println("proceeding with payment")
		p.vm.state = &DispenseState{p.vm}
		return nil
	}
	if amount <= 0 {
		return errors.New("invalid amount inserted")
	}

	return nil
}

func (p *PaymentState) DispenseItem() error {
	return errors.New("no item to take until payment proceeds")
}

func (p *PaymentState) AddItem(n int) error {
	return errors.New("cannot add items while payment proceeds")
}

// DispenseState  represents
type DispenseState struct {
	vm *VendingMachine
}

func (d *DispenseState) RequestItem() error {
	return errors.New("take an item out first")
}

func (d *DispenseState) InsertMoney(amount float64) error {
	return errors.New("take an item out first")
}

func (d *DispenseState) DispenseItem() error {
	d.vm.amount--
	fmt.Println("item dispensed")
	if d.vm.amount == 0 {
		d.vm.state = &NoItemState{d.vm}
		return nil
	}

	d.vm.state = &HasItemState{d.vm}
	return nil
}

func (d *DispenseState) AddItem(n int) error {
	return errors.New("take an item out first")
}

// NoItemState represents
type NoItemState struct {
	vm *VendingMachine
}

func (nis *NoItemState) RequestItem() error {
	return errors.New("no items")
}

func (nis *NoItemState) InsertMoney(amount float64) error {
	return errors.New("no items")
}

func (nis *NoItemState) DispenseItem() error {
	return errors.New("no items")
}

func (nis *NoItemState) AddItem(n int) error {
	nis.vm.amount += n
	nis.vm.state = &HasItemState{nis.vm}
	fmt.Println("items added")
	return nil
}

func main() {
	// init vending machine
	vm := NewVendingMachine(3, 15.0)

	err := vm.RequestItem()
	if err != nil {
		log.Fatalln(err)
	}
	err = vm.InsertMoney(10.0)
	if err != nil {
		log.Fatalln(err)
	}
	err = vm.InsertMoney(5.0)
	if err != nil {
		log.Fatalln(err)
	}
	err = vm.DispenseItem()
	if err != nil {
		log.Fatalln(err)
	}
}

// VendingMachine has a few states
type VendingMachine struct {
	state  State
	amount int
	price  float64
}

func NewVendingMachine(amount int, price float64) *VendingMachine {
	vm := &VendingMachine{nil, amount, price}
	if amount > 0 {
		vm.state = &HasItemState{vm}
	} else {
		vm.state = &NoItemState{vm}
	}

	return vm
}

// RequestItem delegates task to current state object
func (v *VendingMachine) RequestItem() error {
	return v.state.RequestItem()
}

// InsertMoney delegates task to current state object
func (v *VendingMachine) InsertMoney(amount float64) error {
	return v.state.InsertMoney(amount)
}

// DispenseItem delegates task to current state object
func (v *VendingMachine) DispenseItem() error {
	return v.state.DispenseItem()
}

// AddItem delegates task to current state object
func (v *VendingMachine) AddItem(n int) error {
	return v.state.AddItem(n)
}
