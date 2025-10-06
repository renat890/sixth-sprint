package service

import "strings"

func IsMorse(text string) bool { 
	return strings.TrimRight(text, ".- ") == ""
}
	