# FSM-Sim

A simple Finite State Machine Simulation.

## Moore Automata

### Example 1:

**CLI Command:**
```
fsm -s "a,b,c,d" -l "0,1" -t "(a,0,a),(a,1,b),(b,0,a),(b,1,c),(c,0,a),(c,1,d),(d,0,a),(d,1,d)" -m "(a,0),(b,0),(c,0),(d,1)" -c "a" -i "0,1,0,1,0,1,1,1,1,0"
```

**CLI Output:**
```
Input:   0 1 0 1 0 1 1 1 1 0
States: a a b a b a b c d d a
Output: 0 0 0 0 0 0 0 0 1 1 0
```

**Visualization Output:**

![Visualization Output](https://raw.githubusercontent.com/AndreasWillibaldWeber/fsm-sim/main/docs/images/fsm.svg)

### Example 2:

**CLI Command:**
```
fsm -s "a,b,c,d" -a -t "(a,0,a),(a,1,b),(b,0,a),(b,1,c),(c,0,a),(c,1,d),(d,0,a),(d,1,d)" -m "(a,0),(b,0),(c,0),(d,1)" -c "a" -i "0,1,0,1,0,1,1,1,1,0"
```

**CLI Output:**
```
Input:   0 1 0 1 0 1 1 1 1 0
States: a a b a b a b c d d a
Output: 0 0 0 0 0 0 0 0 1 1 0
```

**Visualization Output:**

![Visualization Output](https://raw.githubusercontent.com/AndreasWillibaldWeber/fsm-sim/main/docs/images/fsm_accept.svg)

## Mealy Automata

not yet implemented

## CLI Flags

Use the flag ```-h``` to see the following output:
```
-a    sets the alphabet to {0,1} and visualizes the states mapped to 1 as double circles
-c string
    give a state of states as a starting point as a string
-i string
    give a series of letters as a string 0,1,0,1,1,...
-l string
    give all letters as a string like 0,1,...
-m string
    give a mapping for every state to a letter from the alphabet as a string like (a,0),...
-s string
    give all states as a string like a,b,c,...
-t string
    give all transitions as a string containing triples like (a,0,b),...
```

## License

[MIT](LICENSE)
