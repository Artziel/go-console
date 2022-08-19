package GoConsole

import "time"

func Spinner01() *Spinner {
	chars := "⣾⣽⣻⢿⡿⣟⣯⣷"

	frames := []string{}
	for _, c := range chars {
		frames = append(frames, string(c))
	}

	spinner := Spinner{
		delay:  time.Millisecond * 100,
		Frames: frames,
	}

	return &spinner
}

func Spinner02() *Spinner {
	chars := "⡀⡁⡂⡃⡄⡅⡆⡇⡈⡉⡊⡋⡌⡍⡎⡏⡐⡑⡒⡓⡔⡕⡖⡗⡘⡙⡚⡛⡜⡝⡞⡟⡠⡡⡢⡣⡤⡥⡦⡧⡨⡩⡪⡫⡬⡭⡮⡯⡰⡱⡲⡳⡴⡵⡶⡷⡸⡹⡺⡻⡼⡽⡾⡿⢀⢁⢂⢃⢄⢅⢆⢇⢈⢉⢊⢋⢌⢍⢎⢏⢐⢑⢒⢓⢔⢕⢖⢗⢘⢙⢚⢛⢜⢝⢞⢟⢠⢡⢢⢣⢤⢥⢦⢧⢨⢩⢪⢫⢬⢭⢮⢯⢰⢱⢲⢳⢴⢵⢶⢷⢸⢹⢺⢻⢼⢽⢾⢿⣀⣁⣂⣃⣄⣅⣆⣇⣈⣉⣊⣋⣌⣍⣎⣏⣐⣑⣒⣓⣔⣕⣖⣗⣘⣙⣚⣛⣜⣝⣞⣟⣠⣡⣢⣣⣤⣥⣦⣧⣨⣩⣪⣫⣬⣭⣮⣯⣰⣱⣲⣳⣴⣵⣶⣷⣸⣹⣺⣻⣼⣽⣾⣿"

	frames := []string{}
	for _, c := range chars {
		frames = append(frames, string(c))
	}

	spinner := Spinner{
		delay:  time.Millisecond * 100,
		Frames: frames,
	}

	return &spinner
}

func Spinner03() *Spinner {
	chars := "⋮⋰⋯⋱"

	frames := []string{}
	for _, c := range chars {
		frames = append(frames, string(c))
	}

	spinner := Spinner{
		delay:  time.Millisecond * 100,
		Frames: frames,
	}

	return &spinner
}

func Spinner04() *Spinner {
	chars := "▙▛▜▟"

	frames := []string{}
	for _, c := range chars {
		frames = append(frames, string(c))
	}

	spinner := Spinner{
		delay:  time.Millisecond * 100,
		Frames: frames,
	}

	return &spinner
}

func Spinner05() *Spinner {
	chars := "◢◣◤◥"

	frames := []string{}
	for _, c := range chars {
		frames = append(frames, string(c))
	}

	spinner := Spinner{
		delay:  time.Millisecond * 100,
		Frames: frames,
	}

	return &spinner
}

func Spinner06() *Spinner {
	chars := "◜◝◞◟"

	frames := []string{}
	for _, c := range chars {
		frames = append(frames, string(c))
	}

	spinner := Spinner{
		delay:  time.Millisecond * 100,
		Frames: frames,
	}

	return &spinner
}

func Spinner07() *Spinner {
	chars := "⬖⬘⬗⬙"

	frames := []string{}
	for _, c := range chars {
		frames = append(frames, string(c))
	}

	spinner := Spinner{
		delay:  time.Millisecond * 100,
		Frames: frames,
	}

	return &spinner
}

func Spinner08() *Spinner {
	chars := "⬒⬔⬓⬕"

	frames := []string{}
	for _, c := range chars {
		frames = append(frames, string(c))
	}

	spinner := Spinner{
		delay:  time.Millisecond * 100,
		Frames: frames,
	}

	return &spinner
}

func Spinner09() *Spinner {
	chars := "⯀⯁⯂⬢⯃⯄⬢⯂⯁"

	frames := []string{}
	for _, c := range chars {
		frames = append(frames, string(c))
	}

	spinner := Spinner{
		delay:  time.Millisecond * 100,
		Frames: frames,
	}

	return &spinner
}
