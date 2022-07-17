package main

import "fmt"

// Character type represents player's character.
// Value of this type is constructed via CharacterBuilder
type Character struct {
	Name   string // character's name
	Class  string // character's class
	Race   string // character's race
	Gender string // character's gender
}

func (c Character) String() string {
	return fmt.Sprintf("Name: %s\nClass: %s\nRace: %s\nGender: %s", c.Name, c.Class, c.Race, c.Gender)
}

// CharacterBuilder is an interface for
// builder implementations
type CharacterBuilder interface {
	Gender(g string) CharacterBuilder
	Name(name string) CharacterBuilder
	Class(class string) CharacterBuilder
	Race(race string) CharacterBuilder

	Build() *Character
}

// maleCharacterBuilder is an implementation of
// CharacterBuilder for male characters
type maleCharacterBuilder struct {
	name   string // character's name
	class  string // character's class
	race   string // character's race
	gender string // character's gender
}

// NewMaleCharacterBuilder is a constructor func
func NewMaleCharacterBuilder() *maleCharacterBuilder {
	return &maleCharacterBuilder{gender: "Male"}
}

// Gender sets builder's gender
func (m *maleCharacterBuilder) Gender(g string) CharacterBuilder {
	m.gender = g
	return m
}

// Name sets builder's name
func (m *maleCharacterBuilder) Name(name string) CharacterBuilder {
	m.name = name
	return m
}

// Class sets builder/s class
func (m *maleCharacterBuilder) Class(class string) CharacterBuilder {
	m.class = class
	return m
}

//Race sets builder's race
func (m *maleCharacterBuilder) Race(race string) CharacterBuilder {
	m.race = race
	return m
}

// Build creates Character value and returns a pointer to it
func (m *maleCharacterBuilder) Build() *Character {
	return &Character{
		Name:   m.name,
		Class:  m.class,
		Race:   m.race,
		Gender: m.gender,
	}
}

func main() {
	b := NewMaleCharacterBuilder()
	char := b.Race("Human").Name("Khadgar").Class("Mage").Build()
	fmt.Println(char)
}
