package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	tests := map[string]string{
		"":                 "",
		"a":                "a1",
		"ba":               "a1b1",
		"bca":              "a1b1c1",
		"aaa":              "a3",
		"baa":              "a2b1",
		"bccaa":            "a2b1c2",
		"bbabb":            "a1b4",
		"aabaa":            "a4b1",
		"bbaaa":            "a3b2",
		"aabbaaa":          "a5b2",
		"aaabbbccccc":      "a3b3c5",
		"aaabbbcccccaaaaa": "a8b3c5",
		"zzzzcccUUUuu":     "U3c3u2z4",
		"ЯЯЯБББддд":        "Б3Я3д3"}

	for in, out := range tests {
		if testOut := Counter(in); testOut != out {
			t.Fatalf("Fail. input:%v, out:%v, expected:%v", in, testOut, out)
		}
	}
}
