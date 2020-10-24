package anagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := map[rune]int{}
	for _, v := range s {
		mapS[v]++
	}
	for _, v := range t {
		count, ok := mapS[v]
		if !ok {
			return false
		}
		if count == 1 {
			delete(mapS, v)
		} else {
			mapS[v]--
		}
	}
	return len(mapS) == 0
}
