package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IsMorse(text string) bool { 
	return strings.TrimRight(text, ".- ") == ""
}

func Answer(data string) string {
	var str string
	if IsMorse(data) {
		str = morse.ToText(data)
	} else {
		str = morse.ToMorse(data)
	}
	return str
}