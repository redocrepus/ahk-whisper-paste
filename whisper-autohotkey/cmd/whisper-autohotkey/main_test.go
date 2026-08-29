package main

import "testing"

func boolPointer(value bool) *bool {
	return &value
}

func TestShouldStripNonASCIIEnds(t *testing.T) {
	tests := []struct {
		name       string
		language   string
		configured *bool
		want       bool
	}{
		{name: "English default", language: "en", want: true},
		{name: "English default ignores case", language: "EN", want: true},
		{name: "non-English default", language: "ru", want: false},
		{name: "explicitly enabled for non-English", language: "ru", configured: boolPointer(true), want: true},
		{name: "explicitly disabled for English", language: "en", configured: boolPointer(false), want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := shouldStripNonASCIIEnds(test.language, test.configured); got != test.want {
				t.Fatalf("shouldStripNonASCIIEnds(%q, %v) = %t, want %t", test.language, test.configured, got, test.want)
			}
		})
	}
}

func TestStripNonASCIIEnds(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{name: "removes Unicode garbage from both ends", text: "\u00edhello\u00ed", want: "hello"},
		{name: "removes repeated mixed edge garbage", text: "\n\u00ed hello! \u00ed\r", want: " hello! "},
		{name: "preserves printable ASCII", text: " ~code_123()[]{}! ", want: " ~code_123()[]{}! "},
		{name: "preserves interior Unicode", text: "hello \u00ed world", want: "hello \u00ed world"},
		{name: "removes all garbage", text: "\u00ed\n\u044f", want: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := stripNonASCIIEnds(test.text); got != test.want {
				t.Fatalf("stripNonASCIIEnds(%q) = %q, want %q", test.text, got, test.want)
			}
		})
	}
}
