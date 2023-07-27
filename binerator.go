package binerator

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type generator struct {
	alphabet []any
	timeout  time.Duration
	delay    time.Duration
	output   chan any
	done     chan bool
}

func New(options ...func(*generator)) *generator {
	bi := &generator{
		output: make(chan any),
		done:   make(chan bool),
	}
	for _, option := range options {
		option(bi)
	}
	return bi
}

func WithAlphabet(alphabet ...any) func(*generator) {
	return func(bi *generator) {
		bi.alphabet = append(bi.alphabet, alphabet...)
	}
}

func WithTimeout(timeout time.Duration) func(*generator) {
	return func(bi *generator) {
		bi.timeout = timeout
	}
}

func WithDelay(delay time.Duration) func(*generator) {
	return func(bi *generator) {
		bi.delay = delay
	}
}

func (bi *generator) Emitter() <-chan any {
	if len(bi.alphabet) == 0 {
		fmt.Fprintln(os.Stderr, "Generator needs an alphabet to start emit the sequence.")
		fmt.Fprintln(os.Stderr, "Point out few values for the `WithAlphabet` method.")
		os.Exit(1)
	}
	go func() {
		length := len(bi.alphabet)
		timeout := time.NewTimer(bi.timeout)
		if bi.timeout == 0 {
			timeout.Stop()
		}
		defer func() {
			close(bi.done)
			close(bi.output)
			timeout.Stop()
		}()
		for {
			select {
			case bi.output <- bi.alphabet[rand.Intn(length)]:
				time.Sleep(bi.delay)
			case <-timeout.C:
				return
			case <-bi.done:
				return
			}
		}
	}()
	return bi.output
}

func (bi *generator) Done() {
	bi.done <- true
}
