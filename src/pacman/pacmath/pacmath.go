package pacmath

func Sign(x int) int8 {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}

	return 0
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
