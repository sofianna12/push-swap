# Push-Swap Flowcharts

## push-swap binary flow

```mermaid
flowchart TD
    A([Start]) --> B[Read os.Args]
    B --> C[parser.ParseArgs]
    C --> D{Error?}
    D -->|yes| E[print Error to stderr
os.Exit 1]
    D -->|no| F{len == 0?}
    F -->|yes| G([Exit — no output])
    F -->|no| H[stack.New a with nums
stack.New b empty]
    H --> I[sort.Sort a b stdout]
    I --> J{n?}
    J -->|0 or 1| K[0 ops]
    J -->|2| L[sortTwo
max 1 op]
    J -->|3| M[sortThree
max 2 ops]
    J -->|4 5 6| N[sortSmall
max 11 ops]
    J -->|> 6| O[sortLarge
Turkish algorithm
max 699 ops]
    K --> P([Print ops to stdout
Exit])
    L --> P
    M --> P
    N --> P
    O --> P
```

---

## checker binary flow

```mermaid
flowchart TD
    A([Start]) --> B[Read os.Args]
    B --> C[parser.ParseArgs]
    C --> D{Error?}
    D -->|yes| E[print Error to stderr
os.Exit 1]
    D -->|no| F{len == 0?}
    F -->|yes| G([Exit — no output])
    F -->|no| H[stack.New a with nums
stack.New b empty]
    H --> I[Read line from stdin]
    I --> J{EOF?}
    J -->|no| K[strings.TrimSpace]
    K --> L{empty line?}
    L -->|yes| I
    L -->|no| M[operations.Execute op a b]
    M --> N{known op?}
    N -->|no| O[print Error to stderr
os.Exit 1]
    N -->|yes| I
    J -->|yes| P{a.IsSorted AND
b.Len == 0?}
    P -->|yes| Q([print OK])
    P -->|no| R([print KO])
```

---

## sort algorithm decision tree

```mermaid
flowchart TD
    A[sort.Sort called] --> B{a.Len?}
    B -->|0 or 1| C[return — already sorted]
    B -->|2| D[sortTwo
sa if top > second]
    B -->|3| E[sortThree
hardcoded 5-case
decision tree]
    B -->|4| F[sortSmall
pb min
sortThree
pa]
    B -->|5| G[sortSmall
pb min twice
sortThree
pa pa]
    B -->|6| H[sortSmall
pb min three times
sortThree
pa pa pa]
    B -->|> 6| I[sortLarge
normalise ranks
chunk into sqrt n groups
push chunks to b
pull back to a sorted]
```
