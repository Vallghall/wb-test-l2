package main

import "fmt"

// IGame describes all the games
type IGame interface {
	Name() string
	Rules() string
	Play()
}

// Game defines properties of all the games
type Game struct {
	name  string
	rules string
}

func (g Game) String() string {
	return fmt.Sprintf("Name: \"%s\"\tRules: %s", g.name, g.rules)
}

// BoardGame represents tabletop games
type BoardGame struct {
	Game
}

// NewBoardGame is a constructor func
func NewBoardGame() *BoardGame {
	return &BoardGame{
		Game{
			name:  "Monopoly",
			rules: "*reference to game rules*",
		},
	}
}

// Name is a getter for the game's name
func (bg BoardGame) Name() string {
	return bg.name
}

// Rules is a getter for the games' rules
func (bg BoardGame) Rules() string {
	return bg.rules
}

// Play starts the game
func (bg *BoardGame) Play() {
	fmt.Printf("Let's start playing %s\n", bg.name)
}

// ComputerGame represents PC games
type ComputerGame struct {
	Game
}

// NewComputerGame is a constructor func
// which returns IGame with preordained data
func NewComputerGame() *ComputerGame {
	return &ComputerGame{
		Game{
			name:  "World of Warcraft",
			rules: "No rules, you are on your own",
		},
	}
}

// Name is a getter for the game's name
func (c ComputerGame) Name() string {
	return c.name
}

// Rules is a getter for the game's rules
func (c ComputerGame) Rules() string {
	return c.rules
}

// Play launches the Game
func (c *ComputerGame) Play() {
	fmt.Printf("Launching %s on your PC...", c.name)
}

// GameType defines types of Games
type GameType int

const (
	PC       GameType = iota // ComputerGame type
	Tabletop                 // BoardGame type
)

func GameFactory(gt GameType) IGame {
	switch gt {
	case PC:
		return NewComputerGame()
	case Tabletop:
		return NewBoardGame()
	default:
		panic("invalid game type")
	}
}

func main() {
	computerGame := GameFactory(PC)    // Factory returns ComputerGame as IGame
	boardGame := GameFactory(Tabletop) // Factory returns BoardGame as IGame

	fmt.Printf("%v\n%v\n", computerGame, boardGame)
}
