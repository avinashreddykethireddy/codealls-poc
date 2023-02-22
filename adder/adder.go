package adder

func Add(x, y int) int {
	return x + y
}

func AdderCheckToZero(x, y int) int {
	ans := Add(x, y)
	if ans == 0 {
		return -1
	}
	return 1
}

func CheckValues(x, y int) string {
	sum := x + y
	if sum > 10 {
		return "more"
	} else if sum == 0 {
		return "zero"
	} else {
		return "less"
	}
}
