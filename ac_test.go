package ansicodes

import "testing"

func TestAnsiCodes(t *testing.T) {
	expect := newExpect(t)
	expect(Red, "\033[31m")
	expect(Blue, "\033[34m")
	expect(Bold, "\033[1m")
	expect(YellowBg, "\033[43m")
}

func newExpect(t *testing.T) func(string, string) {
	return func(color string, expected string) {
		if color != expected {
			t.Errorf("%s expected to be %s", color, expected)
		}
	}
}
