package GoConsole

import "testing"

func TestSpinner01(t *testing.T) {

	type testSpinners struct {
		Input    *Spinner
		Expected []string
	}

	tests := []testSpinners{
		{
			Input:    Spinner01(),
			Expected: []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
		},
		{
			Input: Spinner02(),
			Expected: []string{
				"⡀", "⡁", "⡂", "⡃", "⡄", "⡅", "⡆", "⡇", "⡈", "⡉", "⡊", "⡋", "⡌", "⡍", "⡎", "⡏", "⡐", "⡑", "⡒", "⡓", "⡔", "⡕", "⡖", "⡗", "⡘", "⡙", "⡚", "⡛", "⡜", "⡝", "⡞",
				"⡟", "⡠", "⡡", "⡢", "⡣", "⡤", "⡥", "⡦", "⡧", "⡨", "⡩", "⡪", "⡫", "⡬", "⡭", "⡮", "⡯", "⡰", "⡱", "⡲", "⡳", "⡴", "⡵", "⡶", "⡷", "⡸", "⡹", "⡺", "⡻", "⡼", "⡽",
				"⡾", "⡿", "⢀", "⢁", "⢂", "⢃", "⢄", "⢅", "⢆", "⢇", "⢈", "⢉", "⢊", "⢋", "⢌", "⢍", "⢎", "⢏", "⢐", "⢑", "⢒", "⢓", "⢔", "⢕", "⢖", "⢗", "⢘", "⢙", "⢚", "⢛", "⢜",
				"⢝", "⢞", "⢟", "⢠", "⢡", "⢢", "⢣", "⢤", "⢥", "⢦", "⢧", "⢨", "⢩", "⢪", "⢫", "⢬", "⢭", "⢮", "⢯", "⢰", "⢱", "⢲", "⢳", "⢴", "⢵", "⢶", "⢷", "⢸", "⢹", "⢺", "⢻",
				"⢼", "⢽", "⢾", "⢿", "⣀", "⣁", "⣂", "⣃", "⣄", "⣅", "⣆", "⣇", "⣈", "⣉", "⣊", "⣋", "⣌", "⣍", "⣎", "⣏", "⣐", "⣑", "⣒", "⣓", "⣔", "⣕", "⣖", "⣗", "⣘", "⣙", "⣚",
				"⣛", "⣜", "⣝", "⣞", "⣟", "⣠", "⣡", "⣢", "⣣", "⣤", "⣥", "⣦", "⣧", "⣨", "⣩", "⣪", "⣫", "⣬", "⣭", "⣮", "⣯", "⣰", "⣱", "⣲", "⣳", "⣴", "⣵", "⣶", "⣷", "⣸", "⣹",
				"⣺", "⣻", "⣼", "⣽", "⣾", "⣿",
			},
		},
		{
			Input:    Spinner03(),
			Expected: []string{"⋮", "⋰", "⋯", "⋱"},
		},
		{
			Input:    Spinner04(),
			Expected: []string{"▙", "▛", "▜", "▟"},
		},
		{
			Input:    Spinner05(),
			Expected: []string{"◢", "◣", "◤", "◥"},
		},
		{
			Input:    Spinner06(),
			Expected: []string{"◜", "◝", "◞", "◟"},
		},
		{
			Input:    Spinner07(),
			Expected: []string{"⬖", "⬘", "⬗", "⬙"},
		},
		{
			Input:    Spinner08(),
			Expected: []string{"⬒", "⬔", "⬓", "⬕"},
		},
		{
			Input:    Spinner09(),
			Expected: []string{"⯀", "⯁", "⯂", "⬢", "⯃", "⯄", "⬢", "⯂", "⯁"},
		},
	}

	for i, test := range tests {
		if len(test.Input.Frames) != len(test.Expected) {
			t.Errorf("Test Spinner0%d frames count dont match:\ngot  %d\nwant %d", i+1, len(test.Input.Frames), len(test.Expected))
		}

		for j, c := range test.Expected {
			if test.Input.Frames[j] != string(c) {
				t.Errorf("Test Spinner0%d frame %d dont match:\ngot  %s\nwant %s", i+1, j, test.Input.Frames[j], string(c))
			}
		}
	}

}
