type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded_string string
	for _, v := range strs{
		encoded_string += strconv.Itoa(len(v))+"#"+v 
	}
	fmt.Println(encoded_string)
    return encoded_string
}

func (s *Solution) Decode(encoded string) []string {
	strs := []string{}
    j := 0
    for i := 0; i < len(encoded); i++{ 
		if encoded[i] == '#'{
           num, err := strconv.Atoi(string(encoded[j:i]))
           if err != nil{
            fmt.Println(err)
           }
		   strs = append(strs, encoded[i+1:num+i+1])
		   i += num + 1
           j = i
		}
	}
	fmt.Println(strs)
	return strs
}
