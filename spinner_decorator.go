package GoConsole

import (
	"fmt"
	"strings"
	"time"
)

type Spinner struct {
	Frames       []string
	Delay        time.Duration
	currentFrame int
	Prefix       string
	Sufix        string
	stopSpinner  chan struct{}
}

func (s *Spinner) draw() {
	prev := 0
	if s.currentFrame == 0 {
		prev = len(s.Frames) - 1
	}

	fmt.Printf("%s%s%s%s", s.Prefix, strings.Repeat("\r", len(s.Frames[prev])), s.Frames[s.currentFrame], s.Sufix)

	s.currentFrame++
	if s.currentFrame == len(s.Frames) {
		s.currentFrame = 0
	}
}

func (s *Spinner) Stop() {
	s.stopSpinner <- struct{}{}
	close(s.stopSpinner)
}

func (s *Spinner) clear() {
	length := 0
	for _, frame := range s.Frames {
		if len(frame) > length {
			length = len(frame)
		}
	}
	fmt.Printf("\r%s\r", strings.Repeat(" ", len(s.Prefix)+len(s.Sufix)+length))
}

func (s *Spinner) Start() {
	s.stopSpinner = make(chan struct{})
	go func(spinner *Spinner) {
		fmt.Printf("\033[?25l")
		for {
			select {
			case <-s.stopSpinner:
				s.clear()
				fmt.Printf("\033[?25h")
				return
			default:
				s.draw()
				time.Sleep(s.Delay)
			}
		}
	}(s)
}
