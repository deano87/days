package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Pipe represents a variable
type Pipe struct {
	Value   uint64
	Left    string
	Right   string
	Operand string
}

// KnownPipes is a map of all known pipe values
type KnownPipes map[string]uint64

// Pipes represent a map of Pipe
type Pipes map[string]Pipe

// PipesFacade is one structure to group all pipe handlings
type PipesFacade struct {
	Pipes      Pipes
	KnownPipes KnownPipes
}

const noAction = "NONE"

// ParseCommands parses all of the commands given
func (facade *PipesFacade) ParseCommands(commands string) *PipesFacade {
	pipes := Pipes{}
	knownPipes := KnownPipes{}
	commandsArray := strings.Split(commands, "\n")
	for _, command := range commandsArray {
		perform := strings.Split(command, "->")
		perform[0] = strings.Trim(perform[0], " ")
		perform[1] = strings.Trim(perform[1], " ")

		// breakdown
		action := strings.Split(perform[0], " ")
		pipe := Pipe{}

		switch len(action) {
		case 1:
			// add to known pipes
			if isInt(action[0]) {
				knownPipes[perform[1]] = strToUint(action[0])
			} else {
				pipe.Operand = noAction
				pipe.Left = action[0]
				pipes[perform[1]] = pipe
			}
		case 2:
			// unary operand
			pipe.Operand = action[0]
			pipe.Left = action[1]
			pipes[perform[1]] = pipe
		case 3:
			// binary operand
			pipe.Left = action[0]
			pipe.Operand = action[1]
			pipe.Right = action[2]
			pipes[perform[1]] = pipe
		}

	}

	facade.Pipes = pipes
	facade.KnownPipes = knownPipes
	return facade
}

// Resovle calculates the pipes values until nothing is left to calculate
func (facade *PipesFacade) Resovle() *PipesFacade {

	pipes := facade.Pipes
	known := facade.KnownPipes
	for retries := len(pipes); retries > 0; retries-- {
		for pipeName, pipe := range pipes {

			var canCalc = true
			var left uint64
			if isInt(pipe.Left) {
				left = strToUint(pipe.Left)
			} else {
				if val, ok := known[pipe.Left]; ok {
					left = val
					pipe.Left = strconv.FormatUint(left, 16)
				} else {
					canCalc = false
				}
			}

			var right uint64
			if pipe.Operand != "NOT" && pipe.Operand != noAction {
				if isInt(pipe.Right) {
					right = strToUint(pipe.Right)
				} else {
					if val, ok := known[pipe.Right]; ok {
						right = val
						pipe.Right = strconv.FormatUint(right, 16)
					} else {
						canCalc = false
					}
				}
			}

			if canCalc {

				switch pipe.Operand {
				case noAction:
					pipe.Value = left
				case "NOT":
					pipe.Value = 65535 - left
				case "AND":
					pipe.Value = left & right
				case "OR":
					pipe.Value = left | right
				case "LSHIFT":
					pipe.Value = left << right
				case "RSHIFT":
					pipe.Value = left >> right
				}

				facade.KnownPipes[pipeName] = pipe.Value
				delete(pipes, pipeName)
				facade.Pipes = pipes
				if retries < len(pipes) {
					retries = len(pipes)
				}
			}

		}
	}
	return facade
}

// getPipeValue get the value of the pipe
func (facade *PipesFacade) getPipeValue(id string) uint64 {
	if val, ok := facade.KnownPipes[id]; ok {
		return val
	}
	return 0
}

// Print prints all of the pipes
func (facade *PipesFacade) Print() *PipesFacade {
	for pipeName, value := range facade.KnownPipes {
		fmt.Printf("The value of %s is: %d\n", pipeName, value)
	}
	for pipeName := range facade.Pipes {
		fmt.Printf("The value of %s is unknown\n", pipeName)
	}
	return facade
}

func strToUint(s string) uint64 {
	num, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0
	}
	return num
}

func isInt(s string) bool {
	_, err := strconv.ParseUint(s, 10, 0)
	return err == nil
}
