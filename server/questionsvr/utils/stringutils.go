package utils

func IsAnyBlank(strs ...string) bool {
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return true
		}
	}
	return false
}
