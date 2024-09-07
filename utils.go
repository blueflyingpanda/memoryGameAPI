package main

import "regexp"

func IsValidSHA256Hash(s string) bool {
	if len(s) != 64 {
		return false
	}

	hexRegex := regexp.MustCompile("^[0-9a-fA-F]+$")
	return hexRegex.MatchString(s)
}
