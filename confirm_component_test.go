package GoConsole

import "testing"

func TestConfrimComponent(t *testing.T) {

	output := captureOutput(func() {
		confirm := NewConfirm(ConfirmOption{Label: "(Y)es", Value: "y"}, ConfirmOption{Label: "(N)o", Value: "n"})
		confirm.printOptions()
	})

	expected := Colored("<green>(Y)es</green> / <yellow>(N)o</yellow>: ")
	if output != expected {
		t.Errorf("Test ConfirmComponent unexpected output:\ngot  %s\nwant %s", output, expected)
	}

	confirm := NewConfirm(ConfirmOption{Label: "(Y)es", Value: "y"}, ConfirmOption{Label: "(N)o", Value: "n"})
	if _, err := confirm.eval("-"); err == nil {
		t.Errorf("Test ConfirmComponent expected error:\ngot  No Error\nwant %s", ErrConfirmInvalidOption)
	}
	if result, _ := confirm.eval("y"); result != true {
		t.Errorf("Test ConfirmComponent unexpected value:\ngot  %v\nwant TRUE", result)
	}
	if result, _ := confirm.eval("n"); result != false {
		t.Errorf("Test ConfirmComponent unexpected value:\ngot  %v\nwant FALSE", result)
	}
}
