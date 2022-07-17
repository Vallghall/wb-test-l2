package main

import "fmt"

// DamageDealer type represents group damage dealers,
// such as mages, rogues, warriors etc.
type DamageDealer struct {
	// Stats and whatnot
}

// NewDamageDealer is a constructor function
// for DamageDealer type
func NewDamageDealer() *DamageDealer {
	return &DamageDealer{}
}

// DealDamage represents start of damage rotation
func (dd *DamageDealer) DealDamage() {
	fmt.Println("dd does damage rotation")
}

// Support type represents group supports,
// such as holy priests or restoration druids
type Support struct {
	// ...
}

// NewSupport is a constructor function
// for Support type
func NewSupport() *Support {
	return &Support{}
}

// BuffGroup represents support granting buff to their group
func (s *Support) BuffGroup() {
	fmt.Println("sup buffs the team")
}

// HealGroup represents support healing their teammates
func (s *Support) HealGroup() {
	fmt.Println("sup heals the team")
}

// Tank type represents group tanks, such as
// protection paladins and blood death knights
type Tank struct {
	// ...
}

// NewTank is a constructor function
// for Tank type
func NewTank() *Tank {
	return &Tank{}
}

//PullBoss represents Tank initiating a boss fight
func (t *Tank) PullBoss() {
	fmt.Println("tank pulls boss")
}

// DungeonGroup type represents a standard dungeon
// group of 5 people and is a facade for the previous types
type DungeonGroup struct {
	tank    *Tank
	support *Support
	dds     []*DamageDealer
}

// NewDungeonGroup is a constructor function
// for DungeonGroup type
func NewDungeonGroup() *DungeonGroup {
	// DungeonGroup consists of 1 tank, 1 support and 3 damage dealers
	return &DungeonGroup{
		tank:    NewTank(),    // init tank
		support: NewSupport(), // init support
		dds: []*DamageDealer{
			NewDamageDealer(), // init dd
			NewDamageDealer(), // init dd
			NewDamageDealer(), // init dd
		},
	}
}

// FightBoss method simplifies usage of Tank, Support and
// DamageDealer types' methods
func (dg *DungeonGroup) FightBoss(bossName string) {
	dg.support.BuffGroup() // support buffs the team
	dg.tank.PullBoss()     // tank pulls *bossName*
	// every damage dealer starts dealing damage
	for _, dd := range dg.dds {
		dd.DealDamage() // damage rotation
	}
	dg.support.HealGroup() // support heals the team

	fmt.Printf("boss fight with %s goes on!\n", bossName)
}

func main() {
	group := NewDungeonGroup()     // init facade to create a group to go to dungeon
	group.FightBoss("Iron Reaver") // call facade method for boss fight
}
