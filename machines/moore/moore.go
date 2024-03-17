package moore

import (
	"fmt"
	"slices"
)

type Moore struct {
	CurrentState string
	States       []string
	Alphabet     []string
	Transitions  map[string]map[string]string
	Mapping      map[string]string
	Accept       bool
}

func (m *Moore) AddState(state string) {
	if !m.CheckState(state) {
		m.States = append(m.States, state)
	}
}

func (m *Moore) AddLetter(letter string) {
	if !m.CheckLetter(letter) {
		m.Alphabet = append(m.Alphabet, letter)
	}
}

func (m *Moore) AddTransition(state string, input string, nextState string) {
	if m.Transitions == nil {
		m.Transitions = make(map[string]map[string]string)
	}
	if m.Transitions[state] == nil {
		m.Transitions[state] = make(map[string]string)
	}
	m.Transitions[state][input] = nextState
}

func (m *Moore) AddMapping(state string, output string) {
	if m.Mapping == nil {
		m.Mapping = make(map[string]string)
	}
	m.Mapping[state] = output
}

func (m *Moore) StartState(state string) error {
	if !m.CheckState(state) {
		return fmt.Errorf("start (%s) state is not in state set", state)
	}
	m.CurrentState = state
	return nil
}

func (m *Moore) NextState(input string) error {
	if !m.CheckLetter(input) {
		return fmt.Errorf("input (%s) is not in alphabet set", input)
	}
	if m.Transitions[m.CurrentState] == nil || m.Transitions[m.CurrentState][input] == "" {
		return fmt.Errorf("transition (%s -- %s -> ?) is not defined", m.CurrentState, input)
	}
	m.CurrentState = m.Transitions[m.CurrentState][input]
	return nil
}

func (m *Moore) MapOutput() (string, error) {
	if m.Mapping[m.CurrentState] == "" {
		return "", fmt.Errorf("mapping (%s --> ?) is not defined", m.CurrentState)
	}
	return m.Mapping[m.CurrentState], nil
}

func (m *Moore) CheckState(state string) bool {
	return slices.Contains(m.States, state)
}

func (m *Moore) CheckLetter(letter string) bool {
	return slices.Contains(m.Alphabet, letter)
}

func (m *Moore) Validate() bool {
	if len(m.Alphabet) <= 0 || len(m.States) <= 0 {
		return false
	}
	for _, state := range m.States {
		if m.Mapping[state] == "" {
			return false
		}
		if m.Transitions[state] == nil {
			return false
		}
		for _, letter := range m.Alphabet {
			if m.Transitions[state][letter] == "" {
				return false
			}
		}
	}
	return true
}
