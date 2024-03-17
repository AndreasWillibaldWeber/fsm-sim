package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/AndreasWillibaldWeber/fsm-sim/machines"
)

type Flags struct {
	States      string
	Letters     string
	Transitions string
	Mapping     string
	Input       string
	Start       string
	Accept      bool
}

func (f *Flags) RemoveSpaces() {
	f.States = removeAllSpaces(f.States)
	f.Letters = removeAllSpaces(f.Letters)
	f.Transitions = removeAllSpaces(f.Transitions)
	f.Mapping = removeAllSpaces(f.Mapping)
	f.Input = removeAllSpaces(f.Input)
	f.Start = removeAllSpaces(f.Start)
}

func (f *Flags) toConfig() (*machines.Config, error) {
	if f.Accept {
		f.Letters = "0,1"
	}
	transitions, err := parseTriples(f.Transitions)
	if err != nil {
		return nil, err
	}
	mappings, err := parseTuples(f.Mapping)
	if err != nil {
		return nil, err
	}
	return &machines.Config{
		States:      strings.Split(strings.TrimSpace(f.States), ","),
		Letters:     strings.Split(strings.TrimSpace(f.Letters), ","),
		Transitions: transitions,
		Mapping:     mappings,
		Start:       strings.TrimSpace(f.Start),
	}, nil
}

var flags Flags

func init() {

	states := flag.String("s", "", "give all states as a string like a,b,c,...")
	letters := flag.String("l", "", "give all letters as a string like 0,1,...")
	transitions := flag.String("t", "", "give all transitions as a string containing triples like (a,0,b),...")
	mapping := flag.String("m", "", "give a mapping for every state to a letter from the alphabet as a string like (a,0),...")
	input := flag.String("i", "", "give a series of letters as a string 0,1,0,1,1,...")
	start := flag.String("c", "", "give a state of states as a starting point as a string")
	accept := flag.Bool("a", false, "sets the alphabet to {0,1} and visualizes the states mapped to 1 as double circles")

	flag.Parse()

	flags = Flags{
		States:      *states,
		Letters:     *letters,
		Transitions: *transitions,
		Mapping:     *mapping,
		Input:       *input,
		Start:       *start,
		Accept:      *accept,
	}

	flags.RemoveSpaces()
}

func Config() (*machines.Config, error) {
	config, err := flags.toConfig()
	return config, err
}

func Input() []string {
	return strings.Split(strings.TrimSpace(flags.Input), ",")
}

func removeAllSpaces(s string) string {
	return strings.TrimSpace(strings.ReplaceAll(s, " ", ""))
}

func parseTriples(triples string) ([]machines.Transition, error) {
	var transitions []machines.Transition
	for _, triple := range strings.Split(strings.TrimSpace(triples), "),(") {
		transition, err := parseTriple(triple)
		if err != nil {
			return nil, err
		}
		transitions = append(transitions, *transition)
	}

	return transitions, nil
}

func parseTriple(triple string) (*machines.Transition, error) {

	triple = strings.ReplaceAll(triple, "(", "")
	triple = strings.ReplaceAll(triple, ")", "")

	triples := strings.Split(triple, ",")
	if len(triples) < 3 {
		return nil, fmt.Errorf("transition triples does not have the right format")
	}

	return &machines.Transition{
		State:     triples[0],
		Input:     triples[1],
		NextState: triples[2],
	}, nil
}

func parseTuples(tuples string) ([]machines.Mapping, error) {

	var mappings []machines.Mapping
	for _, tuple := range strings.Split(strings.TrimSpace(tuples), "),(") {
		mapping, err := parseTuple(tuple)
		if err != nil {
			return nil, err
		}
		mappings = append(mappings, *mapping)
	}

	return mappings, nil
}

func parseTuple(tuple string) (*machines.Mapping, error) {

	tuple = strings.ReplaceAll(tuple, "(", "")
	tuple = strings.ReplaceAll(tuple, ")", "")

	tuples := strings.Split(tuple, ",")
	if len(tuples) < 2 {
		return nil, fmt.Errorf("mapping tuple does not have the right format")
	}

	return &machines.Mapping{
		State:  tuples[0],
		Output: tuples[1],
	}, nil
}
