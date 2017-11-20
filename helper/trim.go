package helper

import "strings"

// Trim string space
func TrimSpace(str string) string {
	return strings.Trim(str, " ")
}

// Trim string cutset
func TrimCutset(str, cutset string) string {
	return strings.Trim(str, cutset)
}
