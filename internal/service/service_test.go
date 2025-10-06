package service

import (
	"testing"
)

type testCase struct {
	data     string
	expected bool
}

func TestIsMorse(t *testing.T) {
	testData := map[string]testCase{
		"morse": testCase{
			data: ".--. -.-",
			expected: true,
		},
		"text": testCase{
			data: "text",
			expected: false,
		},
	}
	for name, test := range testData {
		t.Run(name, func(t *testing.T) {
			if IsMorse(test.data) != test.expected {
				t.Errorf("Результат обработки IsMorse(%s) не соответствует %t", test.data, test.expected)
			} 
		})
	}
}