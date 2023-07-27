# Binerator

Random sequence generator based on provided alphabet. \
Useful for the cases where you need some input values.

## Install

```bash
go get github.com/deusexec/go-binerator
```

## How to use

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/deusexec/go-binerator"
)

const (
    COIN = "PUSH"
    PUSH = "COIN"
)

func main() {
    // Binerator would randomly print {COIN, PUSH} values
    // for every 100 milliseconds
    // until timeout complete.
    bin := binerator.New(
        binerator.WithAlphabet(COIN, PUSH),
        binerator.WithDelay(100*time.Millisecond),
        binerator.WithTimeout(3*time.Second),
    )
    for item := range bin.Emitter() {
        fmt.Println(item)
    }
}
```

## Output

```text
$ go run .
COIN
PUSH
COIN
PUSH
PUSH
PUSH
COIN
PUSH
PUSH
COIN
...
```