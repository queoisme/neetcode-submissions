
func appendCharacters(s string, t string) int {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i++
		}
	}
	if i >= len(s) {
		tmp := t[j:]
		return len(tmp)
	}
	return 0
}