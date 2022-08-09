package GoConsole

import (
	"testing"
	"time"
)

func SpinnerTest(t *testing.T) {
	spinner := Spinner{}

	spinner.SetDelay(1 * time.Second)
	spinner.SetColor("red")
	spinner.SetPrefix("prefix")
	spinner.SetSufix("sufix")

	if spinner.color != "red" {
		t.Errorf("Test SetColor unexpected color: got %v  want: %v", spinner.color, "red")
	}
	if spinner.prefix != "prefix" {
		t.Errorf("Test SetPrefix unexpected prefix: got %v  want: %v", spinner.prefix, "prefix")
	}
	if spinner.sufix != "sufix" {
		t.Errorf("Test SetSufix unexpected sufix: got %v  want: %v", spinner.sufix, "sufix")
	}
	if spinner.delay != 1*time.Second {
		t.Errorf("Test SetDelay unexpected delay: got %s  want: %v", spinner.delay.String(), "1 second")
	}
}

func Spinner01Test(t *testing.T) {
	spinner := Spinner01()

	spinner.SetDelay(1 * time.Second)
	spinner.SetColor("red")
	spinner.SetPrefix("prefix")
	spinner.SetSufix("sufix")

	if spinner.color != "red" {
		t.Errorf("Test SetColor unexpected color: got %v  want: %v", spinner.color, "red")
	}
	if spinner.prefix != "prefix" {
		t.Errorf("Test SetPrefix unexpected prefix: got %v  want: %v", spinner.prefix, "prefix")
	}
	if spinner.sufix != "sufix" {
		t.Errorf("Test SetSufix unexpected sufix: got %v  want: %v", spinner.sufix, "sufix")
	}
	if spinner.delay != 1*time.Second {
		t.Errorf("Test SetDelay unexpected delay: got %s  want: %v", spinner.delay.String(), "1 second")
	}
}

func Spinner02Test(t *testing.T) {
	spinner := Spinner02()

	spinner.SetDelay(1 * time.Second)
	spinner.SetColor("red")
	spinner.SetPrefix("prefix")
	spinner.SetSufix("sufix")

	if spinner.color != "red" {
		t.Errorf("Test SetColor unexpected color: got %v  want: %v", spinner.color, "red")
	}
	if spinner.prefix != "prefix" {
		t.Errorf("Test SetPrefix unexpected prefix: got %v  want: %v", spinner.prefix, "prefix")
	}
	if spinner.sufix != "sufix" {
		t.Errorf("Test SetSufix unexpected sufix: got %v  want: %v", spinner.sufix, "sufix")
	}
	if spinner.delay != 1*time.Second {
		t.Errorf("Test SetDelay unexpected delay: got %s  want: %v", spinner.delay.String(), "1 second")
	}
}
