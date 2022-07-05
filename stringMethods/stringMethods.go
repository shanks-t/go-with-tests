package stringmethods

import "strings"

func doesContain(str, substr string) bool {
	isASubstring := strings.Contains(str, substr)
	return isASubstring
}
