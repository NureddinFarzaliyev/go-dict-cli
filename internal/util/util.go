package util

import "regexp"

func IsValidTableName(name string) bool {
	valid := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	return valid.MatchString(name)
}
