import (
	"slices"
	"maps"
)

func groupAnagrams(strs []string) [][]string { 
	hash := make(map[string][]string)

    for _, str := range strs{
        temp := []rune(str)
		slices.Sort(temp)
        hash[string(temp)] = append(hash[string(temp)], str)
	}

	return slices.Collect(maps.Values(hash))
}
