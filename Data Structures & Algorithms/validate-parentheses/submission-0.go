func isValid(s string) bool {
    m := map[rune]rune{'}' : '{', ')' : '(', ']' : '['}
	mo := []rune {}
	for _, c := range s{
		dong, ok := m[c]
		if ok{
			if len(mo) == 0{
				return false
			}
			if mo[len(mo) - 1] != dong{
				return false
			}
			index := len(mo) - 1
			mo = mo[:index]
		}else{
			mo = append(mo, c)
		}
	}
	return len(mo) == 0

}
