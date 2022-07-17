package main

import "fmt"

// Command describes command objects
type Command interface {
	Execute()
}

// LightBulb type's methods will be used as commands
type LightBulb struct {
	on bool
}

// On is a LightBulb's state getter
func (lb LightBulb) On() bool {
	return lb.on
}

// SwitchOn switches the LightBulb's state to ON
func (lb *LightBulb) SwitchOn() {
	lb.on = true
}

// SwitchOff switches the LightBulb's state to OFF
func (lb *LightBulb) SwitchOff() {
	lb.on = false
}

// Switch is a command Invoker type
type Switch struct {
	commands map[string]Command
}

func NewSwitch() *Switch {
	return &Switch{
		commands: make(map[string]Command),
	}
}

// Register is invoker's method for memorizing commands
func (s *Switch) Register(name string, cmd Command) {
	s.commands[name] = cmd
}

// Invoke executes a command associated with a given name
func (s *Switch) Invoke(name string) {
	s.commands[name].Execute()
}

// SwitchCommand is a wrapper for LightBulb methods to be invoked
// and implements Command interface
type SwitchCommand struct {
	command func() // wrapped func
}

// Execute calls wrapped function
func (sc SwitchCommand) Execute() {
	sc.command()
}

func main() {
	light := new(LightBulb) // init light bulb
	invoker := NewSwitch()  // init Invoker

	cmdOn := SwitchCommand{light.SwitchOn}   // init Command value with SwitchOn function
	cmdOff := SwitchCommand{light.SwitchOff} // init Command value with SwitchOff function

	invoker.Register("on", cmdOn)   // register command value as "on" command
	invoker.Register("off", cmdOff) // register command value as "off" command

	fmt.Println(light.On()) // print out light bulb state
	invoker.Invoke("on")    // execute command
	fmt.Println(light.On()) // print out light bulb state
	invoker.Invoke("off")   // execute command
	fmt.Println(light.On()) // print out light bulb state
}
