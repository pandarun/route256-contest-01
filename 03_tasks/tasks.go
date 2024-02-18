package main

import (
	"bufio"
	"fmt"
	"io"
)

type Command string

const (
	Start   Command = "M"
	Restart Command = "R"
	Cancel  Command = "C"
	Stop    Command = "D"
)

var StateMachine = map[Command][]Command{
	Start:   {Restart, Cancel, Stop},
	Restart: {Cancel},
	Cancel:  {Start},
	Stop:    {Start},
}

func tasks(input io.Reader, output io.Writer) error {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(input)
	//scanner := bufio.NewScanner(input)
	out = bufio.NewWriter(output)
	defer out.Flush()

	var dataInputs int
	_, err := fmt.Fscan(in, &dataInputs)
	if err != nil {
		return fmt.Errorf("unable to read data inputs")
	}

	var i = 0
	for i < dataInputs {

		var sequence string
		fmt.Fscan(in, &sequence)

		if err != nil {
			return fmt.Errorf("unable to read sequence")
		}

		if len(sequence) < 2 {
			fmt.Fprintln(output, "NO")
			i++
			continue
		}

		first := Command(sequence[0])
		last := Command(sequence[len(sequence)-1])
		if first != Start || last != Stop {
			fmt.Fprintln(output, "NO")
			i++
			continue
		}

		if !validate(sequence) {
			fmt.Fprintln(output, "NO")
		} else {
			fmt.Fprintln(output, "YES")
		}

		i++
	}

	return nil
}

func validate(sequence string) bool {
	for i := 0; i < len(sequence)-1; i++ {

		first := Command(sequence[i])
		next := Command(sequence[i+1])
		if !contains(StateMachine[first], next) {
			return false
		}

	}

	return true

}

func contains(commands []Command, command Command) bool {
	for _, c := range commands {
		if c == command {
			return true
		}
	}
	return false

}
