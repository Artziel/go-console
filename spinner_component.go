package GoConsole

import (
	"fmt"
	"strings"
	"time"
)

type Spinner struct {
	// Public
	prefix string
	sufix  string
	Frames []string

	// Private
	delay        time.Duration
	currentFrame int
	color        string
	stop         chan struct{}
	drawLength   int
}

func (s *Spinner) draw() {
	prev := 0
	if s.currentFrame == 0 {
		prev = len(s.Frames) - 1
	}

	frame := s.Frames[s.currentFrame]
	if s.color != "" {
		frame = Colored(fmt.Sprintf("<"+s.color+">%s</"+s.color+">", s.Frames[s.currentFrame]))
	}

	fmt.Printf("\r%v\r%s%s%s", strings.Repeat(" ", len(s.Frames[prev])), s.prefix, frame, s.sufix)

	s.currentFrame++
	if s.currentFrame == len(s.Frames) {
		s.currentFrame = 0
	}
}

func (s *Spinner) SetDelay(delay time.Duration) *Spinner {
	s.delay = delay

	return s
}

func (s *Spinner) SetPrefix(prefix string) *Spinner {
	s.prefix = prefix

	return s
}

func (s *Spinner) SetSufix(sufix string) *Spinner {
	s.sufix = sufix

	return s
}

func (s *Spinner) SetColor(color string) *Spinner {
	c := strings.ToLower(color)
	s.color = color

	colors := map[string]struct{}{
		"red":    {},
		"green":  {},
		"yellow": {},
		"blue":   {},
		"purple": {},
		"cyan":   {},
		"white":  {},
	}

	if _, found := colors[c]; found {
		s.color = color
	}

	return s
}

func (s *Spinner) Stop() {
	s.stop <- struct{}{}
	close(s.stop)
}

func (s *Spinner) clear() {
	fmt.Printf("\r%s\r", strings.Repeat(" ", s.drawLength))
}

func (s *Spinner) Start() {
	s.stop = make(chan struct{})

	for _, frame := range s.Frames {
		if len(frame) > s.drawLength {
			s.drawLength = len(frame) + len(s.prefix) + len(s.sufix)
		}
	}

	go func(spinner *Spinner) {
		fmt.Printf("\033[?25l")
		for {
			select {
			case <-s.stop:
				s.clear()
				fmt.Printf("\033[?25h")
				return
			default:
				s.draw()
				time.Sleep(s.delay)
			}
		}
	}(s)
}
