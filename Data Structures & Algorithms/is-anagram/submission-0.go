import "slices"
func isAnagram(s string, t string) bool {
	ana1 := []rune(s)
	ana2 := []rune(t)
	slices.Sort(ana1) 
	slices.Sort(ana2)
	return slices.Equal(ana1, ana2)
}
