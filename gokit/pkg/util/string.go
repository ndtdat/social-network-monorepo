package util

func StringInSlice(target string, list []string) bool {
	for _, str := range list {
		if str == target {
			return true
		}
	}

	return false
}
