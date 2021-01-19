package utils

func ContainsInt(target int, source []int) int {
	for i, n := range source {
		if target == n {
			return i
		}
	}
	return -1
}

func ContainsString(target string, source []string) int {
	for i, n := range source {
		if target == n {
			return i
		}
	}
	return -1
}

func ContainsUInt(target uint, source []uint) int {
	for i, n := range source {
		if target == n {
			return i
		}
	}
	return -1
}
