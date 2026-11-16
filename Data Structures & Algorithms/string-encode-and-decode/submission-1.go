type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res []byte

    for _, str := range strs {
        res = append(res, strconv.Itoa(len(str))...)
        res = append(res, '#')
        res = append(res, str...)
    }
	
    return string(res)
}

func (s *Solution) Decode(encoded string) []string {
	  res := []string{}
    i := 0
    for i < len(encoded) {
        length := 0
        for encoded[i] != '#' {
            length = length*10 + int(encoded[i]-'0')
            i++
        }
        i++
        res = append(res, encoded[i:i+length])
        i += length
    }
    return res
}
