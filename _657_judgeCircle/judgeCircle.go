package _657_judgeCircle

func judgeCircle(moves string) bool {

	var h, v int

	for _, c := range moves {
		switch string(c) {
		case "U":
			v++
		case "D":
			v--
		case "L":
			h++
		case "R":
			h--
		default:
			panic("invalid")
		}
	}

	if h == 0 && v == 0 {
		return true
	} else {
		return false
	}
}

