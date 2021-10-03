package bools

func Any(values []bool) bool {
	for _, value := range values {
		if value {
			return true
		}
	}
	return false
}
