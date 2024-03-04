package machines

import (
	"errors"

	"github.com/AndreasWillibaldWeber/fsm-sim/machines/moore"
)

type Transition struct {
	State     string
	Input     string
	NextState string
}

type Mapping struct {
	State  string
	Output string
}

type Config struct {
	States      []string
	Letters     []string
	Transitions []Transition
	Mapping     []Mapping
	Start       string
}

type Result struct {
	Start         string
	Input         []string
	StateSequence []string
	Output        []string
}

func (c *Config) SetupMooreMachine() (*moore.Moore, error) {

	machine := moore.Moore{}

	for _, s := range c.States {
		machine.AddState(s)
	}

	for _, l := range c.Letters {
		machine.AddLetter(l)
	}

	for _, t := range c.Transitions {
		machine.AddTransition(t.State, t.Input, t.NextState)
	}

	for _, m := range c.Mapping {
		machine.AddMapping(m.State, m.Output)
	}

	if err := machine.StartState(c.Start); err != nil {
		return nil, err
	}

	if !machine.Validate() {
		return nil, errors.New("machine has not a valid configuration")
	}

	return &machine, nil
}

func RunMooreMachine(machine *moore.Moore, input []string) (*Result, error) {

	result := Result{
		Start:         machine.CurrentState,
		Input:         input,
		StateSequence: []string{machine.CurrentState},
		Output:        []string{machine.Mapping[machine.CurrentState]},
	}

	for _, l := range input {
		if err := machine.NextState(string(l)); err != nil {
			return nil, err
		}
		result.StateSequence = append(result.StateSequence, machine.CurrentState)
		result.Output = append(result.Output, machine.Mapping[machine.CurrentState])
	}

	return &result, nil
}
