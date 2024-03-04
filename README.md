# FSM-Sim

A simple Finite State Machine Simulation.

## Moore Automata

### Command Line Interface

#### Example

```
fsm -s "a,b,c" -a "0,1" -t "(a,0,b),(a,1,b),(b,0,a),(b,1,c),(c,0,a),(c,1,b)" 
    -m "(a,0),(b,1),(c,1)" -c "a" -i "0,1,0,1,0"
```

#### Flags

Use the flag ```-h``` to see the following output:
```
-a string
    give all letters as a string like 0,1,...
-c string
    give a state of states as a starting point as a string
-i string
    give a series of letters as a string 0,1,0,1,1,...
-m string
    give a mapping for every state to a letter from the alphabet as a string like (a,0),...
-s string
    give all states as a string like a,b,c,...
-t string
    give all transitions as a string containing triples like (a,0,b),...
```

## Mealy Automata

not yet implemented

## License

[MIT](LICENSE)
