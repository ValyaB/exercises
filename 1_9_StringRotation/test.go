package exercises

import "strings"

func isSubString(s1, s2 string) bool {

	return strings.Contains(s1, s2)
}
func stringRotation(s1, s2 string) bool {
	if len(s1) != len(s2) || len(s1) == 0 {
		return false
	}
	s1 = s1 + s1
	return isSubString(s1, s2)
}
