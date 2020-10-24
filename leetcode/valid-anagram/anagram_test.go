package anagram

import "testing"

func TestValid(t *testing.T) {
	s, r := "anagram", "anagram"
	if !IsAnagram(s, r) {
		t.Error("s and r are valid anagram.")
	}
}
func TestInvalid(t *testing.T) {
	s, r := "anagram", "anagra"
	if IsAnagram(s, r) {
		t.Error("s and r are invalid anagram.")
	}
}
func TestRuneValid(t *testing.T) {
	s, r := "hello,世界", "世界,hello"
	if !IsAnagram(s, r) {
		t.Error("s and r are valid anagram.")
	}
}
func TestRuneInValid(t *testing.T) {
	s, r := "hello,宇宙，世界", "世界,hello"
	if IsAnagram(s, r) {
		t.Error("s and r are valid anagram.")
	}
}
