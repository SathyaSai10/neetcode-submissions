func isAnagram(s string, t string) bool {
    anagram := make(map[rune]int, len(s))
	for _, char := range s{
		anagram[char]++
	}
	for _, char1 := range t{
		anagram[char1]--
		if anagram[char1] < 0{
			return false
		}
	}
	for _, char2 := range anagram{
		if char2 != 0{
			return false
		}
	}
	return true
}
