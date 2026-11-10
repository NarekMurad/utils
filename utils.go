package utils

func Contails(a []string, x string) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}

	return false
}