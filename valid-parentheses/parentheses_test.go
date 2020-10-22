package parentheses

import "testing"

func TestIsValidParentheses(t *testing.T) {
	if !IsValid("{{}}(())") {
		t.Error("{{}}(()) are valid parentheses")
	}
	if IsValid("{{}}(()") {
		t.Error("{{}}(() are invalid parentheses")
	}
}
