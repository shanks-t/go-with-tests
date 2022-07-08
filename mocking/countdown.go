package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type DefaultSleeper struct{}
type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration        time.Duration
	ConfiguredSleep func(time.Duration)
}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func (c *ConfigurableSleeper) Sleep() {
	c.ConfiguredSleep(c.Duration)
}
