package utils

func InSlice(a []string, x string) bool {
	for _, s := range a {
		if s == x {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, i := range a {
		if i == x {
			return true
		}
	}
	return false
}
