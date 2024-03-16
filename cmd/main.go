package main

import (
	"fmt"
	"os"
	"strings"

	cli "github.com/AndreasWillibaldWeber/fsm-sim/cmd/cli"
	machine "github.com/AndreasWillibaldWeber/fsm-sim/machines"
	vizualization "github.com/AndreasWillibaldWeber/fsm-sim/machines/vizualization"
	"github.com/goccy/go-graphviz"
)

func main() {

	config, err := cli.Config()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	moore, err := config.SetupMooreMachine()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(2)
	}

	if err := vizualization.DrawMoore(moore, graphviz.SVG, graphviz.CIRCO, "./fsm.svg"); err != nil {
		fmt.Println("error:", err)
		os.Exit(2)
	}

	input := cli.Input()

	result, err := machine.RunMooreMachine(moore, input)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(3)
	}

	fmt.Println("Input:  ", strings.Join(result.Input, " "))
	fmt.Println("States:", strings.Join(result.StateSequence, " "))
	fmt.Println("Output:", strings.Join(result.Output, " "))
}
